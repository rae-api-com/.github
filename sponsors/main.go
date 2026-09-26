// Command sponsors refreshes the sponsors block of a README from GitHub's
// GraphQL API. It only ever rewrites the text between the markers, so the rest
// of the file is untouched and a run with no sponsor changes leaves the file
// byte-identical.
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	markerStart = "<!-- sponsors:start -->"
	markerEnd   = "<!-- sponsors:end -->"

	graphqlURL = "https://api.github.com/graphql"

	// sponsorsQuery asks only for public sponsorships and only for fields a
	// plain token can read: privacyLevel and createdAt would demand the
	// read:user scope, and includePrivate:false already leaves private
	// sponsors out.
	sponsorsQuery = `query($login: String!) {
  organization(login: $login) {
    sponsorshipsAsMaintainer(first: 100, includePrivate: false) {
      nodes {
        sponsorEntity {
          ... on User { login name url avatarUrl }
          ... on Organization { login name url avatarUrl }
        }
      }
    }
  }
}`
)

// sponsor is one public sponsor as rendered in the README.
type sponsor struct {
	Login     string `json:"login"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	AvatarURL string `json:"avatarUrl"`
}

// graphqlResponse mirrors the shape of the sponsorsQuery result.
type graphqlResponse struct {
	Data struct {
		Organization struct {
			SponsorshipsAsMaintainer struct {
				Nodes []struct {
					SponsorEntity sponsor `json:"sponsorEntity"`
				} `json:"nodes"`
			} `json:"sponsorshipsAsMaintainer"`
		} `json:"organization"`
	} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

// client queries GitHub's GraphQL API with a token.
type client struct {
	http  *http.Client
	token string
}

func newClient(token string) *client {
	return &client{http: &http.Client{Timeout: 30 * time.Second}, token: token}
}

// Sponsors returns the organization's public sponsors, in the order GitHub
// reports them.
func (c *client) Sponsors(ctx context.Context, login string) ([]sponsor, error) {
	body, err := json.Marshal(map[string]any{
		"query":     sponsorsQuery,
		"variables": map[string]string{"login": login},
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode query")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, graphqlURL, bytes.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "failed to build request")
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.http.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query sponsors")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.Errorf("sponsors query returned %s", res.Status)
	}

	var payload graphqlResponse
	if err = json.NewDecoder(res.Body).Decode(&payload); err != nil {
		return nil, errors.Wrap(err, "failed to decode sponsors response")
	}
	if len(payload.Errors) > 0 {
		return nil, errors.Errorf("sponsors query failed: %s", payload.Errors[0].Message)
	}

	nodes := payload.Data.Organization.SponsorshipsAsMaintainer.Nodes
	sponsors := make([]sponsor, 0, len(nodes))
	for _, node := range nodes {
		sponsors = append(sponsors, node.SponsorEntity)
	}
	return sponsors, nil
}

// renderSponsors builds one avatar link per sponsor, shaped exactly like the
// hand-written Ko-fi row above it so both read as the same wall of faces. No
// sponsors renders nothing: an empty row beats a sentence that would only be
// right in one of the two languages the profile is written in.
func renderSponsors(sponsors []sponsor) string {
	lines := make([]string, 0, len(sponsors))
	for _, s := range sponsors {
		lines = append(lines, fmt.Sprintf(
			`  <a href="https://github.com/%s"><img src="https://github.com/%s.png" width="48px" alt="@%s" /></a>`,
			s.Login, s.Login, s.Login))
	}
	return strings.Join(lines, "\n")
}

// replaceBlock swaps whatever sits between the markers for block, keeping the
// markers themselves. A README without both markers is an error: silently
// appending the block would bury it wherever the file happened to end.
func replaceBlock(readme, block string) (string, error) {
	start := strings.Index(readme, markerStart)
	if start < 0 {
		return "", errors.Errorf("missing %s marker", markerStart)
	}
	end := strings.Index(readme, markerEnd)
	if end < start {
		return "", errors.Errorf("missing %s marker after %s", markerEnd, markerStart)
	}
	return readme[:start+len(markerStart)] + "\n" + block + "\n" + readme[end:], nil
}

func main() {
	login := flag.String("org", "rae-api-com", "GitHub organization whose sponsors to list")
	path := flag.String("readme", "profile/README.md", "README file holding the sponsors markers")
	flag.Parse()

	token := os.Getenv("SPONSORS_TOKEN")
	if token == "" {
		token = os.Getenv("GITHUB_TOKEN")
	}
	if token == "" {
		fmt.Fprintln(os.Stderr, "SPONSORS_TOKEN or GITHUB_TOKEN is required")
		os.Exit(1)
	}

	sponsors, err := newClient(token).Sponsors(context.Background(), *login)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	current, err := os.ReadFile(*path)
	if err != nil {
		fmt.Fprintln(os.Stderr, errors.Wrap(err, "failed to read README"))
		os.Exit(1)
	}

	updated, err := replaceBlock(string(current), renderSponsors(sponsors))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if updated == string(current) {
		fmt.Printf("sponsors unchanged (%d public)\n", len(sponsors))
		return
	}
	if err = os.WriteFile(*path, []byte(updated), 0o644); err != nil {
		fmt.Fprintln(os.Stderr, errors.Wrap(err, "failed to write README"))
		os.Exit(1)
	}
	fmt.Printf("sponsors updated (%d public)\n", len(sponsors))
}

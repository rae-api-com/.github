package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// newTestSponsor builds a sponsor with every field set, so a test that cares
// about one of them only has to override that one.
func newTestSponsor(login, name string) sponsor {
	return sponsor{
		Login:     login,
		Name:      name,
		URL:       "https://github.com/" + login,
		AvatarURL: "https://avatars.githubusercontent.com/u/1?v=4",
	}
}

func TestRenderSponsors(t *testing.T) {
	t.Run("no sponsors renders an empty block", func(t *testing.T) {
		block := renderSponsors(nil)

		assert.Empty(t, block)
	})

	t.Run("a sponsor renders its profile link", func(t *testing.T) {
		block := renderSponsors([]sponsor{newTestSponsor("calocen", "Carlos Alocen")})

		assert.Contains(t, block, `href="https://github.com/calocen"`)
		assert.Contains(t, block, `alt="@calocen"`)
	})

	t.Run("a sponsor renders the same markup as the hand-written Ko-fi row", func(t *testing.T) {
		block := renderSponsors([]sponsor{newTestSponsor("calocen", "Carlos Alocen")})

		assert.Equal(t,
			`  <a href="https://github.com/calocen"><img src="https://github.com/calocen.png" width="48px" alt="@calocen" /></a>`,
			block)
	})

	t.Run("each sponsor gets its own line", func(t *testing.T) {
		block := renderSponsors([]sponsor{
			newTestSponsor("one", "One"),
			newTestSponsor("two", "Two"),
		})

		assert.Len(t, strings.Split(block, "\n"), 2)
	})
}

func TestReplaceBlock(t *testing.T) {
	const readme = "# Título\n\n" + markerStart + "\nviejo\n" + markerEnd + "\n\n## Licencia\n"

	t.Run("swaps only the text between the markers", func(t *testing.T) {
		updated, err := replaceBlock(readme, "nuevo")

		require.NoError(t, err)
		assert.Equal(t, "# Título\n\n"+markerStart+"\nnuevo\n"+markerEnd+"\n\n## Licencia\n", updated)
	})

	t.Run("replacing with the same block is a no-op", func(t *testing.T) {
		updated, err := replaceBlock(readme, "viejo")

		require.NoError(t, err)
		assert.Equal(t, readme, updated)
	})

	t.Run("a README without the markers is an error", func(t *testing.T) {
		_, err := replaceBlock("# Título\n", "nuevo")

		assert.Error(t, err)
	})

	t.Run("a README missing the closing marker is an error", func(t *testing.T) {
		_, err := replaceBlock("# Título\n"+markerStart+"\n", "nuevo")

		assert.Error(t, err)
	})
}

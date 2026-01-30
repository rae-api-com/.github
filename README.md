# RAE API Organization

Welcome to the **RAE API Organization**. We provide tools and resources for developers who want to interact with the Spanish language programmatically. Our flagship project, **rae-api.com**, is an unofficial API for querying the Real Academia Española (RAE) dictionary.

## Projects

- **[rae-api.com](https://rae-api.com)** — Unofficial API for the RAE dictionary: definitions, conjugations, reverse search, word of the day, polysemic words and locutions.
- **[go-rae](https://github.com/rae-api-com/go-rae)** — Go client library for the RAE API.
- **[rae-tui](https://github.com/rae-api-com/rae-tui)** — Terminal-based interface for exploring the dictionary.
- **[rae-raycast](https://github.com/rae-api-com/rae-raycast)** — Raycast extension for quick lookups.
- **[rae-mcp](https://github.com/rae-api-com/rae-mcp)** — Model Context Protocol server for LLM integration.

## Rate Limiting & API Keys

Due to continuous attacks (DDoS, abusive scraping) affecting service quality, we've implemented rate limiting. **API keys are free** — this isn't monetization, it's abuse protection.

| Tier | Requests/min | Requests/day | How to get it |
|------|-------------|--------------|---------------|
| Free | 10 | 100 | No API key needed |
| Developer | 60 | 5,000 | [Request for free](https://github.com/rae-api-com/.github/issues/new?template=api-key-request.md) |
| Extended | 300 | 50,000 | Request (high-volume projects) |

## In Development

- Fuzzy search for partial matches and derived words
- Inverse indexing for definition-based lookups
- Improved parsing for RAE's HTML changes

## Why this exists

The RAE hasn't provided a public API or database for their dictionary. Many developers have struggled to find a reliable way to query Spanish words programmatically. This project fills that gap.

## FAQ

**Is this official?**  
No. This is an unofficial API with no affiliation to the RAE.

**Can I use this for my project?**  
Yes. It's free for everyone. Donations help keep the servers running: [Ko-Fi](https://ko-fi.com/sonirico).

**What if the RAE changes their website?**  
The API relies on scraping. Changes to their HTML can temporarily break features. We store data locally to ensure continuity.

## Support

If you find this useful, consider supporting the project:
- [Ko-Fi](https://ko-fi.com/sonirico)
- Share rae-api.com with others

Thanks to our supporters:

<p align="center">
  <a href="https://github.com/madebygps"><img src="https://github.com/madebygps.png" width="48px" alt="@madebygps" /></a>
  <a href="https://github.com/nachocerrato"><img src="https://github.com/nachocerrato.png" width="48px" alt="@nachocerrato" /></a>
</p>

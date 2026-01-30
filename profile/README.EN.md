# RAE API Org

Welcome to the **RAE API Organization**. Here you'll find tools for interacting with the Spanish language programmatically. The Real Academia Española (RAE) hasn't shown interest in providing a public API for its dictionary, so we built our own.

This is a non-profit project with educational and collaborative goals. We don't claim ownership over RAE's content. By using this API, you're responsible for how you use its resources.

## Projects

- **[rae-api.com](https://rae-api.com)** — Unofficial API for the RAE dictionary: definitions, conjugations, reverse search, word of the day.
- **[go-rae](https://github.com/rae-api-com/go-rae)** — Go client.
- **[rae-tui](https://github.com/rae-api-com/rae-tui)** — Terminal interface.
- **[rae-raycast](https://github.com/rae-api-com/rae-raycast)** — Raycast extension.
- **[rae-mcp](https://github.com/rae-api-com/rae-mcp)** — LLM integration via Model Context Protocol.

## Rate Limiting & API Keys

Due to continuous attacks (DDoS, abusive scraping) affecting service quality, we've implemented rate limiting. **API keys are free** — this isn't monetization, it's abuse protection.

| Tier | Requests/min | Requests/day | How to get it |
|------|-------------|--------------|---------------|
| Free | 10 | 100 | No API key (anonymous) |
| Developer | 60 | 5,000 | [Request for free](https://github.com/rae-api-com/.github/issues/new?title=[API%20Key%20Request]&labels=api-key-request) |
| Extended | 300 | 50,000 | Request (high-volume projects) |

```bash
curl "https://rae-api.com/api/words/hola?api_key=YOUR_API_KEY"
# or via header
curl -H "X-API-Key: YOUR_API_KEY" "https://rae-api.com/api/words/hola"
```

## In Development

- Fuzzy search for partial matches
- Inverse indexing for definition-based search
- Improved parsing for RAE's HTML changes

## FAQ

**Is this official?**  
No. We have no affiliation with the RAE.

**Can I use this for my project?**  
Yes, it's free to use. Donations welcome: [Ko-Fi](https://ko-fi.com/sonirico).

**What if the RAE changes their website?**  
The API uses scraping. If the structure changes, some features might be temporarily interrupted. We store data locally to ensure continuity.

## Support the project

If you find it useful, consider buying us a coffee at [Ko-Fi](https://ko-fi.com/sonirico) or sharing rae-api.com.

Thanks to those who support this project:

<p align="center">
  <a href="https://github.com/madebygps"><img src="https://github.com/madebygps.png" width="48px" alt="@madebygps" /></a>
  <a href="https://github.com/nachocerrato"><img src="https://github.com/nachocerrato.png" width="48px" alt="@nachocerrato" /></a>
</p>

---

[Leer en español](./README.md)

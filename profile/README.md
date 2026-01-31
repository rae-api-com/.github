# RAE API Org

Bienvenido/a a la **RAE API Organization**. Aquí encontrarás herramientas para interactuar con el idioma español de forma programática. Todo esto surge porque la Real Academia Española parece poco interesada en ofrecer una API pública para su diccionario. Ante la ausencia de soluciones oficiales, hemos decidido aportar la nuestra.

Este proyecto es sin ánimo de lucro, con fines educativos y colaborativos. No pretendemos apropiarnos del contenido de la RAE. Usar esta API depende de ti: cada cual es responsable de cómo emplea estos recursos.

## Proyectos

- **[rae-api.com](https://rae-api.com)** — API no oficial para consultar el diccionario de la RAE: definiciones, conjugaciones, búsquedas inversas, palabra del día.
- **[go-rae](https://github.com/rae-api-com/go-rae)** — Cliente en Go.
- **[rae-tui](https://github.com/rae-api-com/rae-tui)** — Interfaz de terminal.
- **[rae-raycast](https://github.com/rae-api-com/rae-raycast)** — Extensión para Raycast.
- **[rae-mcp](https://github.com/rae-api-com/rae-mcp)** — Integración con LLMs via Model Context Protocol.

## Rate Limiting y API Keys

Debido a continuos ataques (DDoS, scraping abusivo) que afectaban la calidad del servicio, hemos implementado rate limiting. **Las API keys son gratuitas** — esto no es monetización, es protección contra abusos.

| Tier | Requests/min | Requests/día | Cómo obtenerla |
|------|-------------|--------------|----------------|
| Free | 10 | 100 | Sin API key (anónimo) |
| Developer | 60 | 5,000 | [Solicitar gratis](https://github.com/rae-api-com/.github/issues/new?template=api-key-request.md) |
| Extended | 300 | 50,000 | Solicitar (proyectos de alto volumen) |

```bash
curl "https://rae-api.com/api/words/hola?api_key=TU_API_KEY"
# o via header
curl -H "X-API-Key: TU_API_KEY" "https://rae-api.com/api/words/hola"
```

## En Desarrollo

- [x] Definiciones estándar segun dle.rae.es
- [x] Palabra aleatoria con filtros de longitud
- [x] Palabra del día
- [x] Búsqueda difusa para coincidencias parciales
- [x] Índice inverso para búsquedas por definición
- [ ] Mejora de parsing ante cambios del HTML de la RAE

## FAQ

**¿Es esto oficial?**  
No. No hay relación oficial con la RAE.

**¿Puedo usarlo para mi proyecto?**  
Sí, es de uso libre. Donaciones bienvenidas: [Ko-Fi](https://ko-fi.com/sonirico).

**¿Qué pasa si la RAE cambia su web?**  
La API usa scraping. Si cambia la estructura, algunos servicios pueden interrumpirse temporalmente. Almacenamos datos localmente para dar continuidad.

## Apoya el proyecto

Si le sacas partido, considera invitarnos a un café en [Ko-Fi](https://ko-fi.com/sonirico) o compartir rae-api.com.

Gracias a quienes apoyan este proyecto:

<p align="center">
  <a href="https://github.com/madebygps"><img src="https://github.com/madebygps.png" width="48px" alt="@madebygps" /></a>
  <a href="https://github.com/nachocerrato"><img src="https://github.com/nachocerrato.png" width="48px" alt="@nachocerrato" /></a>
</p>

---

[Read in English](./README.EN.md)

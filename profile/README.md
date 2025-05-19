# RAE API Org

Bienvenido/a a la **RAE API Organization**. Aquí encontrarás un conjunto de herramientas y recursos creados para quienes desean interactuar con el idioma español de forma programática. Todo esto surge, en parte, porque la Real Academia Española parece poco interesada en ofrecer una API pública y útil para su diccionario. Así que, ante la ausencia de soluciones oficiales, hemos decidido aportar la nuestra.

Este proyecto es sin ánimo de lucro. Lo hacemos con fines educativos y colaborativos, sin pretender apropiarnos del contenido de la RAE (que, por cierto, mantiene como privado aunque hable del idioma de todos). Usar esta API y sus herramientas depende de ti: cada cual es responsable de cómo emplea estos recursos.

## 🌟 Misión

Queremos facilitar el acceso a los datos del español moderno: definiciones, conjugaciones, y más. Si desarrollas juegos de palabras, creas material educativo, o investigas el lenguaje, puedes encontrar aquí una base útil y sencilla, aunque no oficial.

## 🚀 Proyectos

### 1. **[rae-api.com](https://rae-api.com)**
API no oficial para consultar el diccionario de la RAE. Permite:
- Ver definiciones.
- Revisar conjugaciones verbales.
- Hacer búsquedas inversas (encontrar palabras a partir de una definición).
- Obtener la palabra del día.
- Gestionar palabras polisémicas y locuciones.

### 2. **[go-rae](https://github.com/rae-api-com/go-rae)**
Cliente en Go para integrar fácilmente la RAE API sobre proyectos en este lenguaje.

### 3. **[rae-tui](https://github.com/rae-api-com/rae-tui)**
Interfaz de usuario en la terminal para explorar el diccionario de forma interactiva.

### 4. **[rae-raycast](https://github.com/rae-api-com/rae-raycast)**
Extensión de Raycast para consultar palabras, conjugaciones, palabra del día y obtener palabras aleatorias, directamente desde tu escritorio.

### 5. **[rae-mcp](https://github.com/rae-api-com/rae-mcp)**
La integración con LLMs mediante Model Context Protocol.

## 🛠️ En Desarrollo

- **Búsqueda difusa:** Resultados útiles aun sin coincidencia exacta o con palabras derivadas.
- **Índice inverso:** Para búsquedas por definición.
- **Mejora de parsing:** Adaptándonos a los cambios frecuentes del HTML de la RAE.
- **Autenticación y límites:** Para garantizar un uso razonable y prevenir abusos.
- **Integración con Apify:** Para flujos de trabajo avanzados usando [Apify](https://apify.com/sonirico/diccionario-de-la-real-academia-de-la-lengua-espanola-rae-ppr).

## 💡 Inspiración

La falta de una API oficial —o siquiera una base de datos pública de la RAE— ha hecho que muchos hayamos optado por buscar alternativas: scraping, colaboración y algo de paciencia. Este proyecto pretende, modestamente, llenar ese vacío y facilitar el trabajo a quienes necesitan consultar el diccionario sin complicaciones.

## 🤝 Comunidad

Agradecemos las contribuciones y sugerencias. Si este proyecto te ayuda en tus proyectos (juegos, apps educativas, investigación, integración con sistemas como Anki...), cuéntanoslo y ayuda a mejorar la herramienta.

## 💬 Preguntas Frecuentes

### ¿Es esto oficial?
No. No hay ninguna relación oficial con la RAE. Este proyecto busca facilitar el acceso desarrollando herramientas que la institución, por el momento, no proporciona.

### ¿Puedo usarlo para mi proyecto?
Sí, es de uso libre, aunque cualquier donación para ayudar con los servidores será bienvenida: [BuyMeACoffee](https://buymeacoffee.com/sonirico).

### ¿Qué pasa si la RAE cambia su web?
La API utiliza scraping. Si la estructura de la web cambia, es posible que algunos servicios se interrumpan temporalmente. Aun así, almacenamos datos localmente para ofrecer continuidad en lo posible.

## ❤️ Apoya este proyecto

Dedicamos tiempo y recursos para mantener esta iniciativa. Si le sacas partido, puedes:
- Invitarnos a un café en [BuyMeACoffee](https://buymeacoffee.com/sonirico).
- Compartir **rae-api.com** para que llegue a más gente.

## 📢 Mantente informado/a

Sigue este repositorio para enterarte de novedades, correcciones y futuras mejoras. Siempre estamos abiertos a aportes y ayuda de la comunidad.

Gracias por tu interés y por contribuir a que la lengua española sea un poco más accesible en el mundo digital.

---

[🇬🇧 Read English Version](./README.EN.md)

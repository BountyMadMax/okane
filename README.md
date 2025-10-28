# Okane

## Development

[Echo](https://echo.labstack.com/) - [templ](https://templ.guide/) - [htmx](https://htmx.org/) - [tailwindcss](https://tailwindcss.com/)

### Run server

`air`

or

`templ generate --watch --proxy="http://localhost:8080" --cmd="go run ."`

### Run templ (templates)

`templ generate --watch`

### Run tailwind (CSS)

`npx @tailwindcss/cli -i ./src/assets/css/uncompiled.css -o ./src/assets/css/main.css --watch`


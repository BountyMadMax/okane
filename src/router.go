package app

import (
	"github.com/BountyMadMax/okane/src/controllers"
	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {
	addStaticFiles(e)

	e.GET("/", func(c echo.Context) error {
		return controllers.HomeIndex(c)
	})

	e.GET("/user/new", func(c echo.Context) error {
		return controllers.UserNew(c)
	})
}

func addStaticFiles(e *echo.Echo) {
	e.File("assets/css/main.css", "src/assets/css/main.css")

	e.File("assets/js/htmx.min.js", "node_modules/htmx.org/dist/htmx.min.js")
	e.File("assets/js/alpinejs.min.js", "node_modules/alpinejs/dist/cdn.min.js")
	e.File("assets/js/alpinejs.collapse.min.js", "node_modules/@alpinejs/collapse/dist/cdn.min.js")
	e.File("assets/js/alpinejs.sort.min.js", "node_modules/@alpinejs/sort/dist/cdn.min.js")

	e.File("assets/css/inter.css", "node_modules/@fontsource-variable/inter/index.css")
	e.Static("assets/font/inter/", "node_modules/@fontsource-variable/inter/files/")
}

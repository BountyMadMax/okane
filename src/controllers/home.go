package controllers

import (
	"net/http"

	"github.com/BountyMadMax/okane/src/views"
	"github.com/BountyMadMax/okane/src/views/pages"
	"github.com/labstack/echo/v4"
)

func HomeIndex(c echo.Context) error {
	return views.Render(c, http.StatusOK, pages.Home())
}

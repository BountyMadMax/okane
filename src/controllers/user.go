package controllers

import (
	"net/http"

	"github.com/BountyMadMax/okane/src/database/models"
	"github.com/BountyMadMax/okane/src/views"
	user_pages "github.com/BountyMadMax/okane/src/views/pages/user"
	"github.com/labstack/echo/v4"
)

func UserNew(c echo.Context) error {
	user := models.User{}

	return views.Render(c, http.StatusOK, user_pages.Form(&user))
}

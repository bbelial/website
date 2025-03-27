package router

import (
	"bbelial/template"
	"bbelial/template/static"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Home(ctx echo.Context) error {
	return template.Render(ctx, http.StatusOK, static.Home())
}

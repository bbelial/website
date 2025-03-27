package router

import (
	"bbelial/template"
	"bbelial/template/static"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Project(ctx echo.Context) error {
	return template.Render(ctx, http.StatusOK, static.Project())
}

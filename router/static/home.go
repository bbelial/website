package router

import (
	content "bbelial/router/content"
	"bbelial/template"
	"bbelial/template/static"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Home(pr content.PostReader) echo.HandlerFunc {
	return func(c echo.Context) error {
		posts, err := pr.Query()
		if err != nil {
			return c.HTML(http.StatusInternalServerError, "Errors querying metadata.")
		}

		return template.Render(c, http.StatusOK, static.Home(posts))
	}
}

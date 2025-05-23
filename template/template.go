package template

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Unsafe(html string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, html)
		return
	})
}

func Render(ctx echo.Context, statusCode int, comp templ.Component) error {
	buffer := templ.GetBuffer()
	defer templ.ReleaseBuffer(buffer)

	if err := comp.Render(ctx.Request().Context(), buffer); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buffer.String())
}

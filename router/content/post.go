package router

import (
	"bbelial/model"
	"bbelial/template"
	"bbelial/template/content"
	"bytes"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/labstack/echo/v4"
	"github.com/yuin/goldmark"
)

type PostReader struct {
	Path string
}

func (pr PostReader) Query() ([]model.PostData, error) {
	path := filepath.Join(pr.Path, "*.md")
	fileNames, err := filepath.Glob(path)
	if err != nil {
		return nil, err
	}

	var postsData []model.PostData
	for _, fn := range fileNames {
		var pd model.PostData

		// Opening the files.
		f, err := os.Open(fn)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		// Parsing the metadatas.
		_, err = frontmatter.Parse(f, &pd)
		if err != nil {
			return nil, err
		}

		pd.Slug = strings.TrimSuffix(filepath.Base(fn), ".md")
		postsData = append(postsData, pd)
	}

	return postsData, nil
}

func Article(pr PostReader) echo.HandlerFunc {
	renderer := goldmark.New()
	return func(ctx echo.Context) error {
		var post model.Post

		// Get the slug parameters.
		slug := ctx.Param("slug")

		// Opening the files.
		path := filepath.Join(pr.Path, slug+".md")
		f, err := os.Open(path)
		if err != nil {
			return ctx.HTML(http.StatusInternalServerError, "Errors opening file (possible incorrect path).")
		}
		defer f.Close()

		// Parsing the metadatas.
		md, err := frontmatter.Parse(f, &post)
		if err != nil {
			return ctx.HTML(http.StatusInternalServerError, "Errors parsing metadata.")
		}

		// Rendering the markdown files.
		var buffer bytes.Buffer
		if renderer.Convert(md, &buffer) != nil {
			return ctx.HTML(http.StatusInternalServerError, "Errors rendering content.")
		}

		post.Content = template.Unsafe(buffer.String())

		return template.Render(ctx, http.StatusOK, content.Post(post))
	}
}

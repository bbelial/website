package main

import (
	content "bbelial/router/content"
	router "bbelial/router/static"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Setup the server.
	e := echo.New()
	s := http.Server{
		Addr:    ":2233",
		Handler: e,
	}

	// Serve static files.
	e.Static("/", "public")

	// Initialise posts file reader.
	postReader := content.PostReader{
		Path: "private/article",
	}

	// Routings.
	e.GET("/", router.Home(postReader))
	e.GET("/article", router.Article(postReader))
	e.GET("/article/:slug", content.Article(postReader))
	e.GET("/project", router.Project)

	// Run the server.
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

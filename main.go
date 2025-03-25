package main

import (
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

	// Routings.
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Success!")
	})

	// Run the server.
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

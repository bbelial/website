package main

import (
	"bbelial/router"
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

	// Routings.
	e.GET("/", router.Home)

	// Run the server.
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

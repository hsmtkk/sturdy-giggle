package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	port, err := loadEnvInt("PORT")
	if err != nil {
		log.Fatal(err)
	}
	postgresConfig, err := loadPostgresConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", postgresConfig) // to suppress error

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	h := newHandler()

	// Routes
	e.GET("/", h.Root)
	e.GET("/healthz", h.Healthz)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

type handler struct{}

func newHandler() *handler {
	return &handler{}
}

func (h *handler) Root(ectx echo.Context) error {
	return ectx.String(http.StatusOK, "Root")
}

func (h *handler) Healthz(ectx echo.Context) error {
	return ectx.String(http.StatusOK, "OK")
}

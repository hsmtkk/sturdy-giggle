package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("PORT env var is not defined")
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("failed to parse %s as int: %v", portStr, err)
	}

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

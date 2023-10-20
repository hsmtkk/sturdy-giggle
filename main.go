package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hsmtkk/sturdy-giggle/env"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to init zap logger: %w", err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	reader := env.NewReader()

	port, err := reader.AsInt("PORT")
	if err != nil {
		sugar.Fatal(err)
	}
	postgresConfig, err := reader.PostgresConfig()
	if err != nil {
		sugar.Fatal(err)
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
	if err := e.Start(fmt.Sprintf(":%d", port)); err != nil {
		sugar.Fatal(err)
	}
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

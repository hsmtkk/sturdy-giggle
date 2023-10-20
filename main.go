package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/hsmtkk/sturdy-giggle/env"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	"github.com/jackc/pgx/v5"
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

	ctx := context.Background()

	conn, err := connectPostgres(ctx, postgresConfig)
	if err != nil {
		sugar.Fatal(err)
	}
	defer conn.Close(ctx)

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

func connectPostgres(ctx context.Context, pgConfig env.PostgresConfig) (*pgx.Conn, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", pgConfig.User, pgConfig.Password, pgConfig.Host, pgConfig.Port, pgConfig.Database)
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect Postgres Database: %w", err)
	}
	return conn, nil
}

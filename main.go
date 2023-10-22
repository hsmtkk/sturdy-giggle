package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"text/template"

	"github.com/hsmtkk/sturdy-giggle/env"
	"github.com/hsmtkk/sturdy-giggle/handler"
	"github.com/hsmtkk/sturdy-giggle/repo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to init zap logger: %v", err)
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

	ctx := context.Background()

	conn, err := repo.ConnectPostgres(ctx, postgresConfig)
	if err != nil {
		sugar.Fatal(err)
	}
	defer conn.Close(ctx)

	userRepo := repo.NewUser(conn)
	todoRepo := repo.NewTodo(conn)
	h := handler.New(userRepo, todoRepo)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Templates
	t := &Template{
		templates: template.Must(template.ParseGlob("template/*.html")),
	}
	e.Renderer = t

	// Routes
	e.GET("/", h.Top)
	e.GET("/healthz", h.Healthz)

	// Start server
	if err := e.Start(fmt.Sprintf(":%d", port)); err != nil {
		sugar.Fatal(err)
	}
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

package handler

import (
	"net/http"

	"github.com/hsmtkk/sturdy-giggle/repo"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	userRepo *repo.User
	todoRepo *repo.Todo
}

func New(userRepo *repo.User, todoRepo *repo.Todo) *Handler {
	return &Handler{userRepo, todoRepo}
}

func newHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Top(ectx echo.Context) error {
	return ectx.Render(http.StatusOK, "top", "Der Wille zur Macht")
}

func (h *Handler) Healthz(ectx echo.Context) error {
	return ectx.String(http.StatusOK, "OK")
}

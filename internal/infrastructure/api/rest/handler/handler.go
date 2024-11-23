package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}
func (h *Handler) HelloWold(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"codex-practice/internal/handlers"
)

// Register wires the HTTP routes.
func Register(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})
	e.GET("/health", handlers.Health)
}

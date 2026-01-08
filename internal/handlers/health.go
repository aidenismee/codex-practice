package handlers

import "net/http"

// Health responds with a basic status payload.
func Health(c echoContext) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}

type echoContext interface {
	JSON(code int, i interface{}) error
}

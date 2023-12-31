package middleware

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/rcgc/person-api-echo/authorization"
)

// Authentication .
func Authentication(f echo.HandlerFunc) echo.HandlerFunc{
	return func(c echo.Context) error{
		token := c.Request().Header.Get("Authorization")
		_, err := authorization.ValidateToken(token)
		if err != nil {
			// responder "prohibido"
			return c.JSON(http.StatusForbidden, map[string]string{"error": "no permitido"})
		}

		return f(c)
	}
}
package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"time"
	"url-shortener/internal/pkg"
)

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" || authHeader == "Bearer" {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		jwt := strings.Split(authHeader, " ")[1]
		_, exp, err := pkg.GetIdentity(jwt)

		if int64(exp) < time.Now().Unix() {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		return next(c)
	}
}

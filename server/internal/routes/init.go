package routes

import (
	"github.com/labstack/echo/v4"
	"url-shortener/internal/store"
)

func InitRoutes(app *echo.Echo, store store.InterfaceStore) {
	api := app.Group("/api")

	authControllers(api, store)
	urlControllers(api, store)
}

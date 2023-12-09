package routes

import (
	"github.com/labstack/echo/v4"
	"url-shortener/internal/controllers"
	"url-shortener/internal/store"
)

func authControllers(api *echo.Group, store store.InterfaceStore) {
	controller := controllers.GetAuthControllers(store)

	auth := api.Group("/auth")

	auth.POST("/signup", controller.SignUp)
	auth.POST("/signin", controller.SignIn)
}

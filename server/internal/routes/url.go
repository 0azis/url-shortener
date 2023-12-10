package routes

import (
	"github.com/labstack/echo/v4"
	"url-shortener/internal/controllers"
	"url-shortener/internal/middleware"
	"url-shortener/internal/store"
)

func urlControllers(api *echo.Group, store store.InterfaceStore) {
	controller := controllers.GetUrlControllers(store)

	url := api.Group("/url", middleware.LoggerMiddleware)
	url.POST("", controller.ShortLink)
	url.DELETE("", controller.DeleteLink)
	url.GET("", controller.GetMyUrls)
	api.GET("/:uuid", controller.RedirectByLink)
}

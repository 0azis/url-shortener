package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"url-shortener/internal/routes"
	"url-shortener/internal/store"
)

func InitServer() {
	app := echo.New()

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccessControlAllowOrigin, echo.HeaderAuthorization},
		AllowMethods:     []string{"GET", "HEAD", "PUT", "PATCH", "POST", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}))
	storeInstance := store.GetStoreInstance()
	storeInstance.Open()
	defer storeInstance.Close()

	routes.InitRoutes(app, storeInstance)

	app.Start(":8080")
}

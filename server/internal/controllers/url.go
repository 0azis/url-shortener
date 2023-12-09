package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"url-shortener/internal/models"
	"url-shortener/internal/pkg"
	"url-shortener/internal/services"
	"url-shortener/internal/store"
)

type urlControllers struct {
	UrlServices services.UrlService
}

func (u *urlControllers) ShortLink(c echo.Context) error {
	var newUrlCredentials models.UrlCredentials

	if err := c.Bind(&newUrlCredentials); err != nil {
		return &echo.HTTPError{
			Code:    400,
			Message: "Bad request",
		}
	}

	jwt := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]
	userID, _, _ := pkg.GetIdentity(jwt)

	if !newUrlCredentials.Validate() {
		return &echo.HTTPError{
			Code:    400,
			Message: "Bad request",
		}
	}

	newUrl := models.Url{
		UserID: userID,
		Origin: newUrlCredentials.Origin,
	}

	shortLink, err := u.UrlServices.CreateLink(newUrl)
	if err != nil {
		fmt.Println(err)
		return &echo.HTTPError{
			Code:    500,
			Message: "An error while creating a short link",
		}
	}

	return &echo.HTTPError{
		Code:    201,
		Message: shortLink,
	}
}

func (u *urlControllers) DeleteLink(c echo.Context) error {
	urlID := c.QueryParams().Get("url_id")

	userID, _, _ := pkg.GetIdentity(strings.Split(c.Request().Header.Get("Authorization"), " ")[1])

	err := u.UrlServices.DeleteLink(userID, urlID)

	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
	}

	return &echo.HTTPError{
		Code:    http.StatusOK,
		Message: "An url was deleted",
	}
}

func (u *urlControllers) GetMyUrls(c echo.Context) error {
	userID, _, _ := pkg.GetIdentity(strings.Split(c.Request().Header.Get("Authorization"), " ")[1])

	myUrls, err := u.UrlServices.GetUrls(userID)

	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
	}
	return c.JSON(200, myUrls)
}

func GetUrlControllers(store store.InterfaceStore) urlControllers {
	urlServices := services.GetUrlServices(store)
	return urlControllers{
		UrlServices: &urlServices,
	}
}

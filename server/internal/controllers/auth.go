package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"url-shortener/internal/models"
	"url-shortener/internal/pkg"
	"url-shortener/internal/services"
	"url-shortener/internal/store"
)

type authControllers struct {
	UserServices services.UserService
}

func (a *authControllers) SignUp(c echo.Context) error {
	var newUserCredentials models.Credentials

	if err := c.Bind(&newUserCredentials); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Bad request",
		}
	}

	if err := newUserCredentials.Validate(); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
	}

	newUser := models.User{
		Email:    newUserCredentials.Email,
		Password: newUserCredentials.Password,
	}

	newUserID, err := a.UserServices.CreateUser(newUser)

	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
	}

	token, err := pkg.CreateAccessToken(newUserID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "An error while creating an identification token",
		}
	}

	return &echo.HTTPError{
		Code:    http.StatusOK,
		Message: token,
	}
}

func (a *authControllers) SignIn(c echo.Context) error {
	var credentials models.Credentials

	if err := c.Bind(&credentials); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Bad request",
		}
	}

	fmt.Println(credentials)

	userDB, err := a.UserServices.IsUserExists(credentials.Email)

	if err != nil {
		return &echo.HTTPError{
			Code:    401,
			Message: err,
		}
	}

	if err := pkg.Decode([]byte(userDB.Password), []byte(credentials.Password)); err != nil {
		return &echo.HTTPError{
			Code:    401,
			Message: "Invalid email or password",
		}
	}

	token, err := pkg.CreateAccessToken(userDB.ID)

	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "An error while creating an identification token",
		}
	}

	return &echo.HTTPError{
		Code:    http.StatusOK,
		Message: token,
	}
}

func GetAuthControllers(store store.InterfaceStore) authControllers {
	userServices := services.GetUserServices(store)
	return authControllers{
		UserServices: &userServices,
	}
}

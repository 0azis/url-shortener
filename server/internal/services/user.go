package services

import (
	"errors"
	"url-shortener/internal/models"
	"url-shortener/internal/pkg"
	"url-shortener/internal/store"
)

type UserService interface {
	CreateUser(user models.User) (int, error)
	IsUserExists(email string) (models.User, error)
}

type User struct {
	Store store.InterfaceStore
}

func (u *User) CreateUser(user models.User) (int, error) {

	isValid := u.Store.Users().ValidateEmail(user.Email)

	if isValid {
		return 0, errors.New("Email already exists")
	}

	hashedPassword, err := pkg.Encode([]byte(user.Password))

	if err != nil {
		return 0, errors.New("An error while hashing your password")
	}

	user.Password = string(hashedPassword)

	insertedID, err := u.Store.Users().Insert(user)
	if err != nil {
		return 0, errors.New("An error while creating a user")
	}

	return insertedID, nil
}

func (u *User) IsUserExists(email string) (models.User, error) {
	requestedUser, err := u.Store.Users().GetByEmail(email)

	if err != nil {
		return requestedUser, errors.New("An error while check an user")
	}

	return requestedUser, nil
}

func GetUserServices(store store.InterfaceStore) User {
	return User{Store: store}
}

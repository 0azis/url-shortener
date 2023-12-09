package services

import (
	"errors"
	"url-shortener/internal/models"
	"url-shortener/internal/store"
)

var Domain string = "http://localhost:8080/"

type UrlService interface {
	CreateLink(url models.Url) (models.Url, error)
	DeleteLink(userID int, id string) error
	GetUrls(userID int) ([]models.Url, error)
}

type Url struct {
	Store store.InterfaceStore
}

func (u *Url) CreateLink(url models.Url) (models.Url, error) {
	insertedLink, err := u.Store.Urls().Insert(url)

	if err != nil {
		return insertedLink, errors.New("An error while creating a short link")
	}

	return insertedLink, nil
}

func (u *Url) DeleteLink(userID int, urlID string) error {
	delID, err := u.Store.Urls().Delete(userID, urlID)

	if delID == "" {
		return errors.New("You don't have permission")
	}

	if err != nil {
		return errors.New("An error while deleting an url")
	}

	return nil
}

func (u *Url) GetUrls(userID int) ([]models.Url, error) {

	resultUrls, err := u.Store.Urls().GetByID(userID)

	if err != nil {
		return resultUrls, errors.New("An error while getting your urls")
	}

	return resultUrls, nil
}

func GetUrlServices(store store.InterfaceStore) Url {
	return Url{Store: store}
}

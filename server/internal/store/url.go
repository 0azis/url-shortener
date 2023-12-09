package store

import (
	"database/sql"
	"fmt"
	"url-shortener/internal/models"
	"url-shortener/internal/pkg"
)

var LengthUUID int = 6

type interfaceUrl interface {
	Insert(url models.Url) (models.Url, error)
	Delete(userID int, urlID string) (string, error)
	GetByUserID(userid int) ([]models.Url, error)
	GetOrigin(urlID string) (string, error)
}

type url struct {
	db *sql.DB
}

func (u *url) Insert(url models.Url) (models.Url, error) {
	var insertedUrl models.Url
	row, err := u.db.Query(fmt.Sprintf("insert into urls values ('%s', %d, '%s') returning *", pkg.GenerateUUID(LengthUUID), url.UserID, url.Origin))

	for row.Next() {
		row.Scan(&insertedUrl.ID, &insertedUrl.UserID, &insertedUrl.Origin)
	}

	return insertedUrl, err
}

func (u *url) Delete(userID int, urlID string) (string, error) {
	var deletedID string

	row, err := u.db.Query(fmt.Sprintf("delete from urls where id = '%s' and userid = %d returning id", urlID, userID))

	for row.Next() {
		row.Scan(&deletedID)
	}

	return deletedID, err
}

func (u *url) GetByUserID(userid int) ([]models.Url, error) {
	var resultUrls []models.Url

	rows, err := u.db.Query(fmt.Sprintf("select * from urls where userid = %d", userid))

	if err != nil {
		return resultUrls, err
	}

	for rows.Next() {
		var url models.Url
		rows.Scan(&url.ID, &url.UserID, &url.Origin)
		resultUrls = append(resultUrls, url)
	}

	fmt.Println(resultUrls)

	return resultUrls, nil
}

func (u *url) GetOrigin(urlID string) (string, error) {
	var origin string

	row, err := u.db.Query(fmt.Sprintf("select origin from urls where id = '%s'", urlID))

	if err != nil {
		return "", err
	}

	for row.Next() {
		row.Scan(&origin)
	}

	return origin, nil

}

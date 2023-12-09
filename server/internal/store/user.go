package store

import (
	"database/sql"
	"fmt"
	"url-shortener/internal/models"
)

type interfaceUser interface {
	Insert(user models.User) (int, error)
	ValidateEmail(email string) bool
	GetByEmail(email string) (models.User, error)
}

type user struct {
	db *sql.DB
}

func (u *user) Insert(user models.User) (int, error) {
	var insertedID int

	row, err := u.db.Query(fmt.Sprintf("insert into users (email, password) values ('%s', '%s') returning id", user.Email, user.Password))

	for row.Next() {
		row.Scan(&insertedID)
	}

	return insertedID, err
}

func (u *user) GetByEmail(email string) (models.User, error) {
	var resultUser models.User

	row, err := u.db.Query(fmt.Sprintf("select id, password from users where email = '%s'", email))

	for row.Next() {
		row.Scan(&resultUser.ID, &resultUser.Password)
	}

	return resultUser, err
}

func (u *user) ValidateEmail(email string) bool {
	var validEmail string
	row := u.db.QueryRow(fmt.Sprintf("select email from users where email = '%s'", email))

	row.Scan(&validEmail)

	return validEmail != ""
}

package models

type User struct {
	ID       int    `json:"ID"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

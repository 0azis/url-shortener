package models

import (
	"errors"
	"regexp"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c Credentials) Validate() error {
	isPassword, _ := regexp.Match(`[A-Za-z][0-9]`, []byte(c.Password))

	if isPassword == false || len(c.Password) < 6 || len(c.Password) > 18 {
		return errors.New("Invalid password")
	}

	if c.Email == "" || c.Password == "" {
		return errors.New("Empty data")
	}
	return nil
}

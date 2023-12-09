package pkg

import (
	"github.com/golang-jwt/jwt"
	"time"
)

// KEY Слово-секрет, нужен для расшифровки токена
var KEY = []byte("power")
var TokenTimeAccess int64 = 10000

// CreateAccessToken Метод создания access токена
func CreateAccessToken(userId int) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// Создаем payload структуру
		"userId": userId,
		"exp":    time.Now().Unix() + TokenTimeAccess,
	}).SignedString(KEY)
	return token, err
}

// GetIdentity Расшифровываем токен и получаем из него данные (identity)
func GetIdentity(token string) (int, int, error) {
	identity, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return KEY, nil
	})

	if err != nil {
		return 0, 0, err
	}

	payload := identity.Claims.(jwt.MapClaims)
	userId := int(payload["userId"].(float64))
	exp := int(payload["exp"].(float64))

	return userId, exp, nil

}

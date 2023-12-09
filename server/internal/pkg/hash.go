package pkg

import "golang.org/x/crypto/bcrypt"

func Decode(hash, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err
}

func Encode(password []byte) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return []byte{}, err
	}
	return hashedPassword, nil
}

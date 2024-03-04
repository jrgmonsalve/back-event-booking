package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(password), nil
}

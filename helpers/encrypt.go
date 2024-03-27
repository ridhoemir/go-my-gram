package helpers

import (
	"errors"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (string, error) {
	salt, err := strconv.Atoi(GetEnv("SALT_PASSWORD"))
	if err != nil {
		panic("SALT_PASSWORD must be a number")
	}

	if salt < 4 || salt > 31 {
		err = errors.New("SALT_PASSWORD must be between 4 and 31")
		return "", err
	}
	password := []byte(pass)

	hashed, err := bcrypt.GenerateFromPassword(password, salt)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func ValidateHashPassword(pass, hashedPassword string) bool {
	res := true

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(pass))
	if err != nil {
		res = false
	}

	return res
}

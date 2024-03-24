package helpers

import (
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) string {
	salt, err := strconv.Atoi(GetEnv("SALT_PASSWORD"))
	if err != nil {
		panic("SALT_PASSWORD must be a number")
	}
	password := []byte(pass)

	hashed, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hashed)
}

func ValidateHashPassword(pass, hashedPassword string) bool {
	res := true

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(pass))
	if err != nil {
		res = false
	}

	return res
}

package helpers

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id uint, email string, username string) (string, error) {
	var SECRETKEY = GetEnv("SECRET_KEY")
	claims := jwt.MapClaims{
		"id":       id,
		"email":    email,
		"username": username,
	}

	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	res, err := jwt.SignedString([]byte(SECRETKEY))

	return res, err
}

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	var SECRETKEY = GetEnv("SECRET_KEY")
	err := errors.New("please login to get the token")
	auth := ctx.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(auth, "Bearer ")

	if !bearer {
		return nil, err
	}

	tokenStr := strings.Split(auth, " ")[1]
	if tokenStr == "null" {
		return nil, err
	}
	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			{
				return nil, errors.New("there was an error")
			}
		}
		return []byte(SECRETKEY), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil
}

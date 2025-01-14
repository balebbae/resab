package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretkey = "sdlkfjsdklfjlksdj"

func generateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userID": userId,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString(secretkey)
}
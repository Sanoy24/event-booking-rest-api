package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = "jwtsecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   email,
		"userIid": userId,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

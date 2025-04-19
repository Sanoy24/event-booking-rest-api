package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("jwtsecret")

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	})
	return token.SignedString(secretKey)
}

func VerifyToken(token string) (int64, error) {
	decodedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return 0, fmt.Errorf("token parse error: %w", err)
	}

	if !decodedToken.Valid {
		return 0, errors.New("invalid or expired token")
	}

	claims, ok := decodedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims format")
	}
	fmt.Print(claims)

	// // Safe extraction
	// email, ok := claims["email"].(string)
	// if !ok {
	// 	return errors.New("email claim missing or invalid")
	// }
	// fmt.Println("Email from token:", email)

	userId := int64(claims["userId"].(float64))

	return userId, nil
}

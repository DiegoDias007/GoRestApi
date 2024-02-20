package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(email string, id int) (string, error) {
	userToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    id,
		"exp":   time.Now().Add(time.Hour * 24 * 15).Unix(),
	})
	return userToken.SignedString([]byte(os.Getenv("JWT_KEY")))
}

func ValidateToken(tokenString string) (int, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return 0, false
	}
	userId := getUserId(token)
	return userId, token.Valid
}

func getUserId(token *jwt.Token) int {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0
	}
	userId := int(claims["id"].(float64))
	return userId
}

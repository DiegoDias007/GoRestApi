package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func CreateToken(email string, id int) (string, error)  {
	godotenv.Load()
	userToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"user":  id,
		"exp":   time.Now().Add(time.Hour * 24 * 15).Unix(),
	})
	fmt.Println(os.Getenv("JWT_KEY"))
	return userToken.SignedString([]byte(os.Getenv("JWT_KEY")))
}

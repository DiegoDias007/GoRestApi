package services

import (
	"net/http"

	"api.com/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(
			http.StatusBadRequest, gin.H{"message": "Could not parse data."},
		)
		return
	}
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError, gin.H{"message": "Could not hash password"},
		)
	}
	user.Password = hashedPassword
	err = user.SaveUser()
	if err != nil {
		context.JSON(
			http.StatusInternalServerError, gin.H{"message": "Could not save user."},
		)
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
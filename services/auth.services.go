package services

import (
	"net/http"

	"api.com/models"
	"api.com/utils"
	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	var user *models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(
			http.StatusBadRequest, gin.H{"message": "Could not parse data."},
		)
		return
	}
	err = user.AuthenticateUser()
	if err != nil {
		context.JSON(
			http.StatusUnauthorized, gin.H{"message": "Invalid credentials."},
		)
		return
	}
	token, err := utils.CreateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError, gin.H{"message": "Could not generate token."},
		)
		return
	}
	context.JSON(
		http.StatusOK, gin.H{"message": "Login successful.", "token": token},
	)
}

func SignUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(
			http.StatusBadRequest, gin.H{"message": "Could not parse data."},
		)
		return
	}
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError, gin.H{"message": "Could not hash password"},
		)
	}
	user.Password = hashedPassword
	userId := user.SaveUser()
	if userId == 0 {
		context.JSON(
			http.StatusInternalServerError, gin.H{"message": "Could not save user."},
		)
		return
	}
	token, err := utils.CreateToken(user.Email, userId)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError, gin.H{"message": "Could not generate token."},
		)
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created", "token": token})
}

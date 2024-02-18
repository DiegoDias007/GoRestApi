package controllers

import (
	"api.com/services"
	"github.com/gin-gonic/gin"
)

func RegisterAuthController(server *gin.Engine) {
	server.POST("/signup", services.SignUp)
	server.POST("/login", services.Login)
}
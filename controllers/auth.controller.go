package controllers

import (
	"github.com/gin-gonic/gin"
)

func RegisterAuthController(server *gin.Engine) {
	server.POST("/signup")
}
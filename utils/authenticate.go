package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	userId, isValidToken := ValidateToken(token)
	if !isValidToken {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized user"})
	}
	context.Set("userId", userId)
	context.Next()
}

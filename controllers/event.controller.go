package controllers

import (
	"api.com/services"
	"github.com/gin-gonic/gin"
)

func RegisterEventController(server *gin.Engine) {
	server.GET("/events", services.GetAllEvents)
	server.GET("/events/:eventId", services.GetSingleEvent)
	server.POST("/events", services.AddEvent)
	server.PUT("/events/:eventId", services.UpdateEvent)
}
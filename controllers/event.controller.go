package controllers

import (
	"api.com/services"
	"github.com/gin-gonic/gin"
)

func RegisterEventController(server *gin.Engine) {
	server.GET("/events", services.GetAllEvents)
	server.POST("/events", services.AddEvent)
	server.GET("/events/:eventId", services.GetSingleEvent)
	server.PUT("/events/:eventId", services.UpdateEvent)
	server.DELETE("/events/:eventId", services.DeleteEvent)
}
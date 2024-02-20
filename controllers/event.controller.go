package controllers

import (
	"api.com/services"
	"api.com/utils"
	"github.com/gin-gonic/gin"
)

func RegisterEventController(server *gin.Engine) {
	server.GET("/events", services.GetAllEvents)
	server.POST("/events", utils.Authenticate, services.AddEvent)
	server.GET("/events/:eventId", services.GetSingleEvent)
	server.PUT("/events/:eventId", services.UpdateEvent)
	server.DELETE("/events/:eventId", services.DeleteEvent)
}
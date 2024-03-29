package controllers

import (
	"api.com/services"
	"api.com/utils"
	"github.com/gin-gonic/gin"
)

func RegisterEventController(server *gin.Engine) {
	requireAuth := server.Group("/")
	requireAuth.Use(utils.Authenticate)
	requireAuth.POST("/events", services.AddEvent)
	requireAuth.PUT("/events/:eventId", services.UpdateEvent)
	requireAuth.DELETE("/events/:eventId", services.DeleteEvent)
	requireAuth.POST("/events/:eventId/register", services.RegisterUserForEvent)
	requireAuth.DELETE("/events/:eventId/register", services.DeleteUserRegistration)

	server.GET("/events", services.GetAllEvents)
	server.GET("/events/:eventId", services.GetSingleEvent)
}
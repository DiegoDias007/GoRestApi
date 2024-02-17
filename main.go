package main

import (
	"net/http"
	"time"

	"api.com/database"
	"api.com/models"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	server := gin.Default()
	server.GET("events", getAllEvents)
	server.POST("/events", addEvent)
	server.Run(":8080") // localhost: 8080
}

func addEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	} else {
		event.DateTime = time.Now()
		err := event.SaveEvent()
		if err != nil {
			panic("Error while saving event.")
		}
		context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
	}
}

func getAllEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		panic("Error fetching events.")
	}
	context.JSON(http.StatusOK, gin.H{"message": "Get all events", "events": events})
}
package main

import (
	"net/http"
	"strconv"
	"time"

	"api.com/database"
	"api.com/models"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	server := gin.Default()
	server.GET("/events", getAllEvents)
	server.GET("/events/:eventId", getSingleEvent)
	server.POST("/events", addEvent)
	server.Run(":8080") // localhost: 8080
}

func addEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request, please try again."})
	} else {
		event.DateTime = time.Now()
		err := event.SaveEvent()
		if err != nil {
			context.JSON(
				http.StatusInternalServerError, 
				gin.H{"message": "Error while saving the event, please try again later."},
			)
		}
		context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
	}
}

func getAllEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching events."})
	}
	context.JSON(http.StatusOK, gin.H{"message": "Get all events", "events": events})
}

func getSingleEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching event."})
	}
	event := models.GetSingleEvent(eventId)
	context.JSON(http.StatusOK, gin.H{"message": "Get event", "event": event})
}
package services

import (
	"net/http"
	"strconv"
	"time"

	"api.com/models"
	"github.com/gin-gonic/gin"
)

func AddEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request, please try again."})
		return
	} else {
		event.DateTime = time.Now()
		err := event.SaveEvent()
		if err != nil {
			context.JSON(
				http.StatusInternalServerError,
				gin.H{"message": "Error while saving the event, please try again later."},
			)
			return
		}
		context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
	}
}

func GetAllEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching events."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Get all events", "events": events})
}

func GetSingleEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching event."})
		return
	}
	event, err := models.GetSingleEvent(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Get event", "event": event})
}

func UpdateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing event id"})
		return
	}
	_, err = models.GetSingleEvent(eventId)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"},
		)
		return
	}
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError, gin.H{"message": "Could not parse request data"},
		)
	}
	updatedEvent.ID = int(eventId)
	err = updatedEvent.UpdateEvent()
	if err != nil {
		context.JSON(
			http.StatusInternalServerError, gin.H{"message": "Could not update event."},
		)
	}
	context.JSON(
		http.StatusOK, gin.H{"message": "Event updated", "event": updatedEvent},
	)
}
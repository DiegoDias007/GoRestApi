package models

import (
	"time"

	"api.com/database"
)

type Event struct {
	ID          int
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserID      int
}

func (event *Event) SaveEvent() error {
	var eventId int
	query := `
        INSERT INTO events(name, description, location, dateTime, user_id) 
				VALUES ($1, $2, $3, $4, $5)
				RETURNING id
    `
	err := database.DB.QueryRow(
		query, event.Name, event.Description, event.Location, event.DateTime, event.UserID,
	) .Scan(&eventId)
	
	if err != nil {
		return err
	}
	event.ID = eventId
	return nil
}

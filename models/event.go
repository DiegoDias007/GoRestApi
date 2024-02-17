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

func GetAllEvents() ([]Event, error) {
	var events []Event
	
	query := "SELECT * FROM events"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event Event
		err := rows.Scan(
			&event.ID, &event.Name, &event.Description, &event.Location, 
			&event.DateTime, &event.UserID,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

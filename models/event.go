package models

import (
	"time"

	"api.com/database"
)

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
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
	).Scan(&eventId)

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

func GetSingleEvent(id int64) (*Event, error) {
	var event Event
	query := "SELECT * FROM events WHERE id = $1"
	row := database.DB.QueryRow(query, id)
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location,
		&event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event *Event) UpdateEvent() error {
	query := `
		UPDATE events
		SET name = $1, description = $2, location = $3, dateTime = $4
		WHERE id = $5 
	`
	_, err := database.DB.Exec(
		query, event.Name, event.Description, event.Location, event.DateTime, event.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (event *Event) DeleteEvent() error {
	query := `
		DELETE FROM events WHERE id = $1
	`
	_, err := database.DB.Exec(query, event.ID)
	if err != nil {
		return err
	}
	return nil
}

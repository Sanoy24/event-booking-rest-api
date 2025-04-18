package models

import (
	"time"

	databse "github.com/sanoy24/event-booking-rest-api/database"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	query := `
	INSERT INTO events(name,description,location,date_time,user_id)
	VALUES (?,?,?,?,?)
	`
	stmt, err := databse.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}
	id, err := result.LastInsertId()

	// e.ID := id
	e.ID = id

	return err

	// will be added to db

}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := databse.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)

	}
	return events, nil
}

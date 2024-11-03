package models

import (
	"REST_PROJECT/db"
	"time"
)

// Event represents an event with its details
type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int64     `json:"user_id"`
}

// Save stores the event into the database
func (e *Event) Save() error {
	query :=
		`INSERT INTO events (name, description, location, dateTime, user_id)
    	 VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId() // Get the ID of the newly inserted row
	e.ID = id                        // Update the ID field of the struct

	return err
}

// GetAllEvents retrieves all events from the database
func GetAllEvents() ([]Event, error) {
	query :=
		`SELECT * FROM events `

	// Execute the query and get the rows
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	// Iterate over the rows, creating a new Event for each and adding it to the slice of events
	for rows.Next() {
		var event Event

		// Scan the columns of the row into the fields of the Event struct and add it to the slice of events
		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, err
}

// GetEventByID retrieves a single event by its ID from the database
func GetEventByID(id int64) (*Event, error) {
	query :=
		`SELECT * FROM events WHERE id = ?`

	// Execute the query and get the row
	// and use the id parameter to fill in the placeholder in the query
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

// Update modifies an existing event in the database
func (e Event) Update() error {
	query :=
		`UPDATE events
	     SET name = ?, description = ?, location = ?, dateTime = ?
	     WHERE id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)

	return err
}

// Delete removes an existing event from the database
func (e Event) Delete() error {
	query :=
		`DELETE FROM events
   		 WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)

	return err
}

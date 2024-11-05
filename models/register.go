package models

import "REST_PROJECT/db"

// Register adds a new registration for a user to an event
// The registration is stored in the database and linked to the user and event
func (e Event) Register(userId int64) error {
	query :=
		`INSERT INTO registrations (event_id, user_id) VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

// CancelRegistration removes a user's registration from an event
func (e Event) CancelRegistration(userId int64) error {

	query :=
		`DELETE FROM registrations WHERE event_id = ? AND user_id = ? `

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

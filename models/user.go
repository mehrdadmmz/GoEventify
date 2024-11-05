package models

import (
	"REST_PROJECT/db"
	"REST_PROJECT/utils"
	"errors"
)

// User represents a user with email and password
type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required""`
	Password string `json:"password" binding:"required"`
}

// Save stores the user into the database with hashed password
func (u User) Save() error {
	query :=
		`INSERT INTO users (email, password)
   		 VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId() // Get the ID of the newly inserted row
	u.ID = userId                        // Update the ID field of the struct

	return err
}

// ValidateCredentials checks if the provided credentials match the ones in the database
func (u *User) ValidateCredentials() error {
	query :=
		`SELECT id, password FROM users WHERE email = ?`

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword) // Scan the retrieved password into the variable
	if err != nil {
		return errors.New("credentials invalid")
	}

	// Check if the provided password matches the retrieved password
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}

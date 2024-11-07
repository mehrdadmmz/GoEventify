package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes the provided password using bcrypt and returns the hashed password as a string.
func HashPassword(password string) (string, error) {

	// GenerateFromPassword returns the hashed password. It returns an error if the password is empty or if the cost is less than the minimum allowed cost (4).
	// costs define the algorithmic cost that should be used. The work factor is 2**cost, and the default cost is 10
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares the provided password with the hashed password and returns true if they match.
func CheckPasswordHash(password, hashedPassword string) bool {

	// CompareHashAndPassword compares a hashed password with its possible plaintext equivalent. Returns nil on success, or an error on failure.
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

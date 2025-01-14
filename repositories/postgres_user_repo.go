package repositories

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
)

type userRepoImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepoImpl{db: db}
}

func (repo *userRepoImpl) ValidateUser(username, password string) (bool, error) {
	var count int
	hashedPassword := hashPassword(password) // Ensure this function hashes using SHA-256

	query := "SELECT COUNT(*) FROM users WHERE username = $1 AND password_hash = $2"
	row := repo.db.QueryRow(query, username, hashedPassword)

	if err := row.Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}

func (repo *userRepoImpl) CreateUser(username, password string) (bool, error) {
	hashedPassword := hashPassword(password) // Ensure this function hashes using SHA-256
	query := "INSERT INTO users (username, password_hash) VALUES ($1, $2)"
	_, err := repo.db.Exec(query, username, hashedPassword)
	if err != nil {
		return false, err
	}
	return true, nil
}

func hashPassword(password string) string {
	// Implement the logic to hash the password using SHA-256, for example:
	return fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
}

func (repo *userRepoImpl) DeleteUser(username string) (bool, error) {
	query := "DELETE FROM users WHERE username = $1"
	_, err := repo.db.Exec(query, username)
	if err != nil {
		return false, err
	}
	return true, nil
}

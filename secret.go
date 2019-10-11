package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Secret struct {
	ID         int
	Ciphertext string
	UserID     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

var db *sql.DB

func InitConnection() {
	connStr := "user=postgres dbname=secret sslmode=disable host=localhost"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

// GetSecret returns a secret with secretID
func GetSecret(secretID int) *Secret {
	InitConnection()
	rows, err := db.Query("SELECT id, ciphertext, user_id, created_at, updated_at, deleted_at FROM secrets WHERE id = $1", secretID)
	if err != nil {
		log.Printf("Secret select failed: %s", err)
		return nil
	}
	defer rows.Close()

	var secret Secret
	for rows.Next() {
		secret = Secret{}
		rows.Scan(&secret.ID, &secret.Ciphertext, &secret.UserID, &secret.CreatedAt, &secret.UpdatedAt, &secret.DeletedAt)
	}

	return &secret
}

func CreateSecret(secret *Secret) error {
	InitConnection()
	_, err := db.Exec("INSERT INTO secrets(id, ciphertext, user_id, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6)", secret.ID, secret.Ciphertext, secret.UserID, secret.CreatedAt, secret.UpdatedAt, secret.DeletedAt)
	if err != nil {
		return errors.Wrap(err, "Secret insert failed")
	}

	return nil
}

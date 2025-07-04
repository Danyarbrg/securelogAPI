package db

import (
	"context"
	"database/sql"
	"log"
	"time"
	"fmt"

	_"github.com/jackc/pgx/v5/stdlib" // мы импортируем только init() функцию внутри пакета.
	"github.com/Danyarbrg/securelogAPI/internal/models"

)

func Connect(databaseURL string) *sql.DB {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatalf("Error when openning a connection with DB: %v\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("Failed to connect to the database: %v\n", err)
	}

	fmt.Println("Successful conection!")
	return db
}

func InsertLog(db *sql.DB, entry models.LogEntry) error {
	query := `
		INSERT INTO logs (event_type, user_id, payload)
		VALUES ($1, $2, $3)
	`
	_, err := db.Exec(query, entry.EventType, entry.UserID, entry.Payload)
	if err != nil {
		return err
	}

	fmt.Println("Log successfully written!")
	return nil
}
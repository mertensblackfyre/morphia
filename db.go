package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func run() (err error) {

	url := TURSO_DATABASE_URL + "?authToken=" + TURSO_AUTH_TOKEN

	// Open database connection
	db, err := sql.Open("libsql", url)
	if err != nil {
		return fmt.Errorf("error opening cloud db: %w", err)
	}
	defer db.Close()

	// Configure connection pool
	db.SetConnMaxIdleTime(9 * time.Second)

	ctx := context.Background()

	schema := `
	CREATE TABLE IF NOT EXISTS roles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE,
		color TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		role TEXT,
		booster INTEGER,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (role) REFERENCES roles(name)
			ON UPDATE CASCADE
			ON DELETE SET NULL
	);
	`

	// Create test table
	_, err = db.ExecContext(ctx, schema)
	if err != nil {
		return fmt.Errorf("error creating table: %w", err)
	}

	return nil
}
func InsertRoleDB(db *sql.DB,role_id string , name string) {
	_, err := db.Exec("INSERT INTO role (id, name) VALUES (?, ?)", role_id, name)
	if err != nil {
		fmt.Errorf("error inserting data: %w", err)
	}
	fmt.Println("Inserted test data")
}

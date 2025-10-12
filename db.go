package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

// declare globally
var DB *sql.DB

func Init() (err error) {

	url := TURSO_DATABASE_URL + "?authToken=" + TURSO_AUTH_TOKEN

	// Open database connection
	DB, err = sql.Open("libsql", url)
	if err != nil {
		return fmt.Errorf("error opening cloud DB: %w", err)
	}

	// Configure connection pool
	DB.SetConnMaxIdleTime(9 * time.Second)

	ctx := context.Background()

	schema := `
	PRAGMA foreign_keys = ON;


	CREATE TABLE IF NOT EXISTS roles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id TEXT,
		role_id TEXT,
		name TEXT,
		color TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP);
	`

	// Create test table
	_, err = DB.ExecContext(ctx, schema)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return nil
}
func InsertRoleDB(role_id string, name string, user_id string, color string) {
	_, err := DB.ExecContext(context.Background(),"INSERT INTO roles (role_id,name, user_id, color) VALUES (?, ?,?,?)", role_id, name, user_id, color)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Inserted test data")
}

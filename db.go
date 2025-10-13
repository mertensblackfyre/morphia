package main

import (
	"context"
	"database/sql"
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
		Sugar.Errorw("error opening cloud DB: %w", err)
		return err
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
		Sugar.Errorw("Error creating table", "error", err)
		return err
	}

	return nil
}
func InsertRoleDB(role_id string, name string, user_id string, color string) {
	_, err := DB.ExecContext(context.Background(), "INSERT INTO roles (role_id,name, user_id, color) VALUES (?,?,?,?)", role_id, name, user_id, color)
	if err != nil {
		Sugar.Errorw("Error inserting data", "error", err)
		return
	}
	Sugar.Info("Successfully inserted to the database")
}

func RemoveRoleDB(role_id string, user_id string) {
	_, err := DB.ExecContext(context.Background(), "DELETE FROM roles WHERE user_id = ? AND role_id = ?", user_id, role_id)
	if err != nil {
		Sugar.Errorw("Error removing data", "error", err)
		return
	}

	Sugar.Info("Successfully deleted from the databases")
}

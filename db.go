package main

import (
	"context"
	"database/sql"
	"errors"
	"time"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

/*

	CREATE TABLE IF NOT EXISTS  server_boosters(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id TEXT UNIQUE,
		role_id TEXT,
		name TEXT,
		username TEST,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP);
	
*/

var DB *sql.DB

func Init() (err error) {

	url := TURSO_DATABASE_URL + "?authToken=" + TURSO_AUTH_TOKEN

	DB, err = sql.Open("libsql", url)
	if err != nil {
		Sugar.Errorw("error opening cloud DB: %w", err)
		return err
	}

	DB.SetConnMaxIdleTime(9 * time.Second)

	ctx := context.Background()

	schema := `
	PRAGMA foreign_keys = ON;

	CREATE TABLE IF NOT EXISTS roles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id TEXT UNIQUE,
		role_id TEXT,
		name TEXT,
		color TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP);
	`

	_, err = DB.ExecContext(ctx, schema)
	if err != nil {
		Sugar.Errorw("Error creating table", "error", err)
		return err
	}

	return nil
}

func UpdateRoleDB(name string, color string, role_id string) {
	res, err := DB.ExecContext(context.Background(), "UPDATE roles SET name = ?, color = ? WHERE role_id = ?", name, color, role_id)
	if err != nil {
		Sugar.Errorw("failed to update role",
			"error", err,
			"role_id", role_id,
		)
		return
	}

	rowsAffected, _ := res.RowsAffected()
	Sugar.Infow("role updated",
		"role_id", role_id,
		"rows_affected", rowsAffected,
	)
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

	Sugar.Info("Successfully deleted from the database")
}

func TraverseDB(user_id string) *Role {
	rows, err := DB.QueryContext(context.Background(), "SELECT * from roles")

	if err != nil {
		Sugar.Errorln(err)
	}

	var role Role
	for rows.Next() {
		if err := rows.Scan(&role.ID, &role.UserID, &role.RoleID, &role.Name, &role.Color, &role.CreatedAt); err != nil {
			Sugar.Errorln("Error scanning row:", err)
			return nil
		}
		if role.UserID == user_id {
			if err != nil {
				Sugar.Errorln(err)
				return nil
			}
		} else {
			Sugar.Errorw("Error: Role not found", "error", err)
			return nil

		}
	}
	return &role
}

func CheckUserHasRole(user_id string) int {

	var r Role
	query := `SELECT id, user_id, role_id, name, color, created_at FROM roles WHERE user_id = ?`
	err := DB.QueryRowContext(context.Background(), query, user_id).Scan(
		&r.ID, &r.UserID, &r.RoleID, &r.Name, &r.Color, &r.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			Sugar.Infow("User does not have a role", "user_id", user_id)
			return 1
		} else {
			Sugar.Errorw("Query failed", "error", err)
			return 2
		}
	}

	Sugar.Infow("User has a role")
	return 0
}

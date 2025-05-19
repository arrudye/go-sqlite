package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectSQLite(path string) (*sql.DB, error) {
	if _, err := os.Stat(filepath.Dir(path)); os.IsNotExist(err) {
		return nil, fmt.Errorf("Data folder doesn't exist: %s", filepath.Dir(path))
	}

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("Connection error: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Connection check failed: %w", err)
	}

	log.Printf("DB created/connected: %s", path)
	return db, nil
}

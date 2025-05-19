package database

import (
	"database/sql"
	"log"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(db *sql.DB, migrationsPath string) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return err
	}

	cleanPath := filepath.ToSlash(migrationsPath)
	if !filepath.IsAbs(cleanPath) {
		cleanPath = "/" + cleanPath
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+cleanPath,
		"sqlite3",
		driver,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	log.Println("Migrations completed successfully")
	return nil
}

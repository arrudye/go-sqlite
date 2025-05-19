package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	DataDir       string
	MigrationsDir string
	DBPath        string
}

func NewConfig(projectRoot, dataDir, migrationsDir, dbFileName string) *Config {
	return &Config{
		DataDir:       filepath.Join(projectRoot, dataDir),
		MigrationsDir: filepath.Join(projectRoot, migrationsDir),
		DBPath:        filepath.Join(dataDir, dbFileName),
	}
}

func InitializeDB(config *Config) (*sql.DB, error) {
	if err := os.MkdirAll(config.DataDir, 0755); err != nil {
		log.Println("Failed to create data folder: ", err)
		return nil, err
	}

	db, err := ConnectSQLite(config.DBPath)
	if err != nil {
		log.Println("Failed to connect to the database: ", err)
		return nil, err
	}

	if err := RunMigrations(db, config.MigrationsDir); err != nil {
		log.Println("Failed to run the migrations: ", err)
		db.Close()
		return nil, err
	}

	return db, nil
}

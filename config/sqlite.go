package config

import (
	"os"

	"github.com/MarcosGomesDev/goopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/main.db"

	// Check if the databse file exists
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Infof("database file does not exist, creating it")

		// Create the database file and directory
		err := os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			logger.Errorf("error creating database directory: %v", err)
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			logger.Errorf("error creating database file: %v", err)
			return nil, err
		}

		file.Close()
	}

	// Create DB and connect
	db, err := gorm.Open(sqlite.Open("./db/main.db"), &gorm.Config{})

	if err != nil {
		logger.Errorf("error initializing sqlite: %v", err)
		return nil, err
	}

	// Migrate Schema
	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("error migrating schemas: %v", err)
		return nil, err
	}

	// Return the DB
	return db, nil
}

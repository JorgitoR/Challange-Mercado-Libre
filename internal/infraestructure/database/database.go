package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

func NewDatabase() (*gorm.DB, error) {

	log.Info("Setting up new database connection")

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")
	sslmode := os.Getenv("SSL_MODE")

	connectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", dbHost, dbPort, dbUsername, dbTable, dbPassword, sslmode)

	db, err := gorm.Open("postgres", connectString)
	if err != nil {
		return db, err
	}

	if err := db.DB().ping(); err != nil {
		return db, err
	}

	return db, nil
}

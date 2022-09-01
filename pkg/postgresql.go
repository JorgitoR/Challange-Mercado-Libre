package pkg

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

func PostgresClient() (*gorm.DB, error) {
	log.Info("Setting up new database connection")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := "5432"

	connectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)
	fmt.Println("Conexion dtb", connectString)

	db, err := gorm.Open("postgres", connectString)
	if err != nil {
		return db, err
	}

	if err := db.DB().Ping(); err != nil {
		return db, err
	}
	return db, nil
}

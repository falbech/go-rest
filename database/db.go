package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDB() {
	connectionString := "host=localhost user=root password=root port=5430 dbname=root sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Error connecting to database!")
	}
}

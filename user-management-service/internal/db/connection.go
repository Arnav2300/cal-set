package db

import (
	"log"
	"os"

	"github.com/Arnav2300/cal-set/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabase() *gorm.DB {
	dsn := os.Getenv("POSTGRESQL_URI")
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	sqldb, err := connection.DB()
	if err != nil {
		log.Fatalf("Failed to get generic database object: %v", err)
	}
	err = sqldb.Ping()
	if err != nil {
		log.Fatalln("database connected!")
	}
	log.Println("connected to database!")
	return connection
}

func InitialMigration() {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.AutoMigrate(models.User{})
}

func CloseDatabase(connection *gorm.DB) {
	sqldb, err := connection.DB()
	if err != nil {
		log.Println("failed to close connection to databse!")
	}
	sqldb.Close()
}

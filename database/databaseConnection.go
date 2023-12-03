package database

import (
	"fmt"
	"log"
	"os"

	"github.com/resturant/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	// Create or open the log file
	logFile, err := os.OpenFile("./logs/database.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error opening log file: ", err)
	}
	defer logFile.Close()

	// Set log output to the log file
	log.SetOutput(logFile)

	// Log messages will now be written to the file
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db

	// Check if the database exists
	var count int64
	if err := DB.Raw("SELECT COUNT(*) FROM pg_database WHERE datname = ?", "resturant").Scan(&count).Error; err != nil {
		log.Fatal("Error checking if the database exists: ", err)
	}

	// If the database doesn't exist, create it
	if count == 0 {
		if err := DB.Exec("CREATE DATABASE resturant").Error; err != nil {
			log.Fatal("Error creating database: ", err)
		}
		fmt.Println("Database 'resturant' created")
	}

	// Connect to the created database
	db, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db

	// AutoMigrate should be called after establishing the connection
	err = DB.AutoMigrate(&models.Customer{}, &models.Order{}, &models.OrderItem{}, &models.Payment{}, &models.Staff{})
	if err != nil {
		log.Fatal("Error auto migrating models: ", err)
	}

	fmt.Println("Connected to PostgreSQL and migrated models")
}

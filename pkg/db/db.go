package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB initializes the database connection using PostgreSQL.
func ConnectDB() {
	var err error
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	if dbname == "" {
		log.Println("Warning: DB_NAME is not set, defaulting to 'postgres'")
		dbname = "postgres"
	}

	sslmode := os.Getenv("DB_SSLMODE")
	if sslmode == "" {
		sslmode = "disable"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s search_path=public client_encoding=UTF8",
		host, user, password, dbname, port, sslmode,
	)

	fmt.Printf("Connecting to database: %s on %s:%s\n", dbname, host, port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Ensure 'public' schema exists (sometimes missing in new databases or special environments)
	if err := DB.Exec("CREATE SCHEMA IF NOT EXISTS public").Error; err != nil {
		log.Printf("Warning: failed to ensure public schema exists: %v", err)
	}

	fmt.Println("Database connection successfully opened")
}

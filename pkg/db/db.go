package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB initializes the database connection using PostgreSQL.
func ConnectDB() {
	var err error
	var dsn string

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL != "" {
		dsn = databaseURL
	} else {
		host := os.Getenv("DB_HOST")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")
		port := os.Getenv("DB_PORT")
		sslmode := os.Getenv("DB_SSLMODE")

		if dbname == "" {
			log.Println("Warning: DB_NAME is not set, defaulting to 'postgres'")
			dbname = "postgres"
		}

		if sslmode == "" {
			sslmode = "disable"
		}

		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s search_path=public client_encoding=UTF8",
			host, user, password, dbname, port, sslmode,
		)
	}

	fmt.Println("Connecting to database...")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Connection pooling configuration (Optimized for Serverless)
	sqlDB, err := DB.DB()
	if err == nil {
		// Low limits for serverless to prevent 'too many connections'
		sqlDB.SetMaxIdleConns(2)
		sqlDB.SetMaxOpenConns(5)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

	// Ensure 'public' schema exists
	if err := DB.Exec("CREATE SCHEMA IF NOT EXISTS public").Error; err != nil {
		log.Printf("Warning: failed to ensure public schema exists: %v", err)
	}

	fmt.Println("Database connection successfully opened")
}

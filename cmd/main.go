package main

import (
	"log"
	"os"
	"tasklybe/internal/db"
	"tasklybe/internal/siswa"
	"tasklybe/internal/task"
	"tasklybe/internal/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

// @title           Taskly API
// @version         1.0
// @description     A simple task management API.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.email  fiber@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host            localhost:8080
// @BasePath        /api
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Connect to the database
	db.ConnectDB()

	// Auto-migrate models
	err := db.DB.AutoMigrate(&user.User{}, &task.Task{}, &siswa.Siswa{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize Fiber app
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	// Initialize services
	userService := user.NewService(db.DB)
	taskService := task.NewService(db.DB)
	siswaService := siswa.NewService(db.DB)

	// Initialize handlers
	userHandler := user.NewHandler(userService)
	taskHandler := task.NewHandler(taskService)
	siswaHandler := siswa.NewHandler(siswaService)

	// Setup routing
	api := app.Group("/api")
	user.SetupUserRoutes(api, userHandler)
	task.SetupTaskRoutes(api, taskHandler)
	siswa.SetupSiswaRoutes(api, siswaHandler)

	// Start server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}

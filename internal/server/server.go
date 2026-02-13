package server

import (
	"log"
	"tasklybe/internal/db"
	"tasklybe/internal/siswa"
	"tasklybe/internal/task"
	"tasklybe/internal/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func SetupApp() *fiber.App {
	// Load .env file (it's okay if it fails on Vercel as env vars are set in dashboard)
	_ = godotenv.Load()

	// Connect to the database
	db.ConnectDB()

	// Auto-migrate models
	err := db.DB.AutoMigrate(&user.User{}, &task.Task{}, &siswa.Siswa{})
	if err != nil {
		log.Printf("Failed to migrate database: %v", err)
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

	return app
}

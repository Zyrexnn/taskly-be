package server

import (
	"log"
	"tasklybe/pkg/db"
	"tasklybe/pkg/siswa"
	"tasklybe/pkg/task"
	"tasklybe/pkg/user"

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
		log.Println("Database migration error (continuing):", err)
	} else {
		log.Println("Database migration completed successfully")

		// Seed default user if not exists (helpful for first-time Vercel deploy)
		var count int64
		db.DB.Model(&user.User{}).Count(&count)
		if count == 0 {
			log.Println("No users found, seeding default user 'ikhsan'...")
			userService := user.NewService(db.DB)
			_, _ = userService.Register(user.RegisterRequestDTO{
				Name:     "ikhsan",
				Email:    "ikhsan@example.com",
				Password: "123",
			})
		}
	}

	// Initialize Fiber app
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}))

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

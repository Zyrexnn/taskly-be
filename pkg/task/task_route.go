package task

import (
	"tasklybe/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupTaskRoutes(router fiber.Router, handler *Handler) {
	taskGroup := router.Group("/tasks", middleware.Protected()) // Apply JWT middleware here

	taskGroup.Post("/", handler.CreateTask)
	taskGroup.Get("/", handler.GetAllTasks)
	taskGroup.Get("/:id", handler.GetTaskByID)
	taskGroup.Put("/:id", handler.UpdateTask)
	taskGroup.Delete("/:id", handler.DeleteTask)
}

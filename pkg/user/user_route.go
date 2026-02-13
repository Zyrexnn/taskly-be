package user

import "github.com/gofiber/fiber/v2"

func SetupUserRoutes(router fiber.Router, handler *Handler) {
	userGroup := router.Group("/user")
	userGroup.Post("/register", handler.Register)
	userGroup.Post("/login", handler.Login)
}

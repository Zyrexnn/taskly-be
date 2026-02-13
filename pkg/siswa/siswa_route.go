package siswa

import "github.com/gofiber/fiber/v2"

func SetupSiswaRoutes(router fiber.Router, handler *Handler) {
	siswaGroup := router.Group("/siswa")
	siswaGroup.Post("/", handler.Create)
	siswaGroup.Get("/", handler.GetAll)
	siswaGroup.Get("/:id", handler.GetByID)
	siswaGroup.Put("/:id", handler.Update)
	siswaGroup.Delete("/:id", handler.Delete)
}

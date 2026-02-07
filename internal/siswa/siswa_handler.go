package siswa

import (
	"strconv"
	"tasklybe/internal/dto"
	"tasklybe/internal/validation"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// Create godoc
// @Summary      Create a new siswa
// @Description  Create a new student record
// @Tags         Siswa
// @Accept       json
// @Produce      json
// @Param        siswa  body      CreateSiswaRequestDTO  true  "Siswa data"
// @Success      201    {object}  dto.ResponseWrapper[SiswaResponseDTO]
// @Failure      400    {object}  dto.ResponseWrapper[any]
// @Failure      500    {object}  dto.ResponseWrapper[any]
// @Router       /siswa [post]
func (h *Handler) Create(c *fiber.Ctx) error {
	var req CreateSiswaRequestDTO
	if ok, errors := validation.BindAndValidate(c, &req); !ok {
		return c.Status(fiber.StatusBadRequest).JSON(dto.NewErrorResponse("Validation failed", errors))
	}

	siswa, err := h.service.Create(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.NewErrorResponse("Gagal membuat data siswa", err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(dto.NewSuccessResponse(siswa, "Siswa berhasil ditambahkan"))
}

// GetAll godoc
// @Summary      Get all siswa
// @Description  Retrieve all students with pagination and search
// @Tags         Siswa
// @Accept       json
// @Produce      json
// @Param        page    query     int     false  "Page number"        default(1)
// @Param        limit   query     int     false  "Items per page"     default(10)
// @Param        search  query     string  false  "Search by nama or NIS"
// @Success      200     {object}  dto.ResponseWrapper[SiswaListResponseDTO]
// @Failure      500     {object}  dto.ResponseWrapper[any]
// @Router       /siswa [get]
func (h *Handler) GetAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	search := c.Query("search", "")

	result, err := h.service.GetAll(page, limit, search)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.NewErrorResponse("Gagal mengambil data siswa", err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.NewSuccessResponse(result, "Data siswa berhasil diambil"))
}

// GetByID godoc
// @Summary      Get siswa by ID
// @Description  Retrieve a student by their ID
// @Tags         Siswa
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Siswa ID"
// @Success      200  {object}  dto.ResponseWrapper[SiswaResponseDTO]
// @Failure      404  {object}  dto.ResponseWrapper[any]
// @Router       /siswa/{id} [get]
func (h *Handler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.NewErrorResponse("ID tidak valid", err.Error()))
	}

	siswa, err := h.service.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.NewErrorResponse("Siswa tidak ditemukan", err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.NewSuccessResponse(siswa, "Data siswa ditemukan"))
}

// Update godoc
// @Summary      Update siswa
// @Description  Update a student's information
// @Tags         Siswa
// @Accept       json
// @Produce      json
// @Param        id     path      int                    true  "Siswa ID"
// @Param        siswa  body      UpdateSiswaRequestDTO  true  "Updated siswa data"
// @Success      200    {object}  dto.ResponseWrapper[SiswaResponseDTO]
// @Failure      400    {object}  dto.ResponseWrapper[any]
// @Failure      404    {object}  dto.ResponseWrapper[any]
// @Router       /siswa/{id} [put]
func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.NewErrorResponse("ID tidak valid", err.Error()))
	}

	var req UpdateSiswaRequestDTO
	if ok, errors := validation.BindAndValidate(c, &req); !ok {
		return c.Status(fiber.StatusBadRequest).JSON(dto.NewErrorResponse("Validation failed", errors))
	}

	siswa, err := h.service.Update(uint(id), req)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.NewErrorResponse("Gagal mengupdate siswa", err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.NewSuccessResponse(siswa, "Siswa berhasil diupdate"))
}

// Delete godoc
// @Summary      Delete siswa
// @Description  Delete a student by their ID
// @Tags         Siswa
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Siswa ID"
// @Success      200  {object}  dto.ResponseWrapper[any]
// @Failure      404  {object}  dto.ResponseWrapper[any]
// @Router       /siswa/{id} [delete]
func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.NewErrorResponse("ID tidak valid", err.Error()))
	}

	if err := h.service.Delete(uint(id)); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.NewErrorResponse("Gagal menghapus siswa", err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.NewSuccessResponse[any](nil, "Siswa berhasil dihapus"))
}

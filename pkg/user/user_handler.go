package user

import (
	"tasklybe/pkg/dto"
	"tasklybe/pkg/validation"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// Register godoc
// @Summary      Register a new user
// @Description  Create a new user account
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user  body      RegisterRequestDTO  true  "User registration data"
// @Success      201  {object}  dto.ResponseWrapper[UserResponseDTO]
// @Failure      400  {object}  dto.ResponseWrapper[any]
// @Failure      500  {object}  dto.ResponseWrapper[any]
// @Router       /user/register [post]
func (h *Handler) Register(c *fiber.Ctx) error {
	var req RegisterRequestDTO
	if ok, errors := validation.BindAndValidate(c, &req); !ok {
		return c.Status(fiber.StatusBadRequest).JSON(dto.NewErrorResponse("Validation failed", errors))
	}

	user, err := h.service.Register(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.NewErrorResponse("Failed to register user", err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(dto.NewSuccessResponse(user, "User registered successfully"))
}

// Login godoc
// @Summary      Login a user
// @Description  Authenticate a user and get a JWT
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        credentials  body      LoginRequestDTO  true  "User login credentials"
// @Success      200          {object}  dto.ResponseWrapper[LoginResponseDTO]
// @Failure      400          {object}  dto.ResponseWrapper[any]
// @Failure      401          {object}  dto.ResponseWrapper[any]
// @Router       /user/login [post]
func (h *Handler) Login(c *fiber.Ctx) error {
	var req LoginRequestDTO
	if ok, errors := validation.BindAndValidate(c, &req); !ok {
		return c.Status(fiber.StatusBadRequest).JSON(dto.NewErrorResponse("Validation failed", errors))
	}

	loginData, err := h.service.Login(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.NewErrorResponse("Login failed", err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.NewSuccessResponse(loginData, "Login successful"))
}

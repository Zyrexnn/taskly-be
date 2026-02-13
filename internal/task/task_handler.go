package task

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

func (h *Handler) getUserIDFromLocals(c *fiber.Ctx) (uint, error) {
	id, ok := c.Locals("userId").(uint)
	if !ok {
		return 0, fiber.NewError(fiber.StatusUnauthorized, "Cannot parse user ID")
	}
	return id, nil
}

// CreateTask godoc
// @Summary      Create a new task
// @Description  Add a new task for the logged-in user
// @Tags         Task
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        task  body      CreateTaskDTO  true  "Task creation data"
// @Success      201  {object}  dto.ResponseWrapper[Task]
// @Failure      400  {object}  dto.ResponseWrapper[any]
// @Failure      401  {object}  dto.ResponseWrapper[any]
// @Router       /tasks [post]
func (h *Handler) CreateTask(c *fiber.Ctx) error {
	userID, err := h.getUserIDFromLocals(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.NewErrorResponse(err.Error(), nil))
	}

	var req CreateTaskDTO
	if ok, errors := validation.BindAndValidate(c, &req); !ok {
		return c.Status(fiber.StatusBadRequest).JSON(dto.NewErrorResponse("Validation failed", errors))
	}

	task, err := h.service.CreateTask(userID, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.NewErrorResponse("Failed to create task", err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(dto.NewSuccessResponse(task, "Task created successfully"))
}

// GetAllTasks godoc
// @Summary      Get all tasks
// @Description  Get all tasks for the logged-in user
// @Tags         Task
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {object}  dto.ResponseWrapper[[]Task]
// @Failure      401  {object}  dto.ResponseWrapper[any]
// @Router       /tasks [get]
func (h *Handler) GetAllTasks(c *fiber.Ctx) error {
	userID, err := h.getUserIDFromLocals(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.NewErrorResponse(err.Error(), nil))
	}

	tasks, err := h.service.GetAllTasks(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.NewErrorResponse("Failed to retrieve tasks", err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.NewSuccessResponse(&tasks, "Tasks retrieved successfully"))
}

// GetTaskByID godoc
// @Summary      Get a single task
// @Description  Get a single task by its ID
// @Tags         Task
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id   path      int  true  "Task ID"
// @Success      200  {object}  dto.ResponseWrapper[Task]
// @Failure      401  {object}  dto.ResponseWrapper[any]
// @Failure      404  {object}  dto.ResponseWrapper[any]
// @Router       /tasks/{id} [get]
func (h *Handler) GetTaskByID(c *fiber.Ctx) error {
	userID, err := h.getUserIDFromLocals(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.NewErrorResponse(err.Error(), nil))
	}

	taskID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.NewErrorResponse("Invalid task ID", nil))
	}

	task, err := h.service.GetTaskByID(userID, uint(taskID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.NewErrorResponse("Task not found", err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.NewSuccessResponse(task, "Task retrieved successfully"))
}

// UpdateTask godoc
// @Summary      Update a task
// @Description  Update a task by its ID
// @Tags         Task
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id   path      int  true  "Task ID"
// @Param        task body      UpdateTaskDTO  true  "Task update data"
// @Success      200  {object}  dto.ResponseWrapper[Task]
// @Failure      400  {object}  dto.ResponseWrapper[any]
// @Failure      401  {object}  dto.ResponseWrapper[any]
// @Failure      404  {object}  dto.ResponseWrapper[any]
// @Router       /tasks/{id} [put]
func (h *Handler) UpdateTask(c *fiber.Ctx) error {
	userID, err := h.getUserIDFromLocals(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.NewErrorResponse(err.Error(), nil))
	}

	taskID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.NewErrorResponse("Invalid task ID", nil))
	}

	var req UpdateTaskDTO
	if ok, errors := validation.BindAndValidate(c, &req); !ok {
		return c.Status(fiber.StatusBadRequest).JSON(dto.NewErrorResponse("Validation failed", errors))
	}

	task, err := h.service.UpdateTask(userID, uint(taskID), req)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.NewErrorResponse("Failed to update task", err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.NewSuccessResponse(task, "Task updated successfully"))
}

// DeleteTask godoc
// @Summary      Delete a task
// @Description  Delete a task by its ID
// @Tags         Task
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id   path      int  true  "Task ID"
// @Success      200  {object}  dto.ResponseWrapper[any]
// @Failure      401  {object}  dto.ResponseWrapper[any]
// @Failure      404  {object}  dto.ResponseWrapper[any]
// @Router       /tasks/{id} [delete]
func (h *Handler) DeleteTask(c *fiber.Ctx) error {
	userID, err := h.getUserIDFromLocals(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.NewErrorResponse(err.Error(), nil))
	}

	taskID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.NewErrorResponse("Invalid task ID", nil))
	}

	err = h.service.DeleteTask(userID, uint(taskID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.NewErrorResponse("Failed to delete task", err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.NewSuccessResponse[any](nil, "Task deleted successfully"))
}

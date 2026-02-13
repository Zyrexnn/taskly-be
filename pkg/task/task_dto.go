package task

// CreateTaskDTO defines the structure for creating a new task.
type CreateTaskDTO struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}

// UpdateTaskDTO defines the structure for updating an existing task.
// Pointers are used to distinguish between a field not being provided
// and a field being set to its zero value (e.g., Completed: false).
type UpdateTaskDTO struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Completed   *bool   `json:"completed"`
}

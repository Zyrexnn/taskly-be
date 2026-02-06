package dto

// ResponseWrapper is a generic struct for standard API responses.
type ResponseWrapper[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    *T     `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}

// NewSuccessResponse creates a success response with data.
func NewSuccessResponse[T any](data *T, message string) ResponseWrapper[T] {
	return ResponseWrapper[T]{
		Success: true,
		Message: message,
		Data:    data,
	}
}

// NewErrorResponse creates an error response.
func NewErrorResponse(message string, err any) ResponseWrapper[any] {
	return ResponseWrapper[any]{
		Success: false,
		Message: message,
		Error:   err,
	}
}
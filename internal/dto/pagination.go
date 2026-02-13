package dto

import "math"

// PaginatedResponse is a generic struct for paginated data.
type PaginatedResponse[T any] struct {
	Data       []T   `json:"data"`
	Total      int64 `json:"total"`
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalPages int   `json:"total_pages"`
}

// NewPaginatedResponse creates a new paginated response.
func NewPaginatedResponse[T any](data []T, total int64, page, limit int) PaginatedResponse[T] {
	totalPages := 0
	if limit > 0 {
		totalPages = int(math.Ceil(float64(total) / float64(limit)))
	}

	return PaginatedResponse[T]{
		Data:       data,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}
}

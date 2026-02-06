package user

import (
	"tasklybe/internal/task"
	"time"

	"gorm.io/gorm"
)

// User represents the user model.
type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Email     string         `gorm:"unique;not null" json:"email"`
	Password  string         `gorm:"not null" json:"-"` // Omit from JSON responses
	Tasks     []task.Task    `gorm:"foreignKey:UserID" json:"tasks,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

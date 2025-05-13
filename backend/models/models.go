package models

import (
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description,omitempty"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var ValidStatuses = []string{"new", "in_progress", "done"}

func ValidateStatus(status string) bool {
	for _, valid := range ValidStatuses {
		if status == valid {
			return true
		}
	}
	return false
}

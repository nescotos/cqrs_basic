package model

import "time"

// Tweet : Defines a Tweet Schema configuration
type Tweet struct {
	ID        string    `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"` // To perform soft delete
}

package model

import (
	"time"

	"github.com/google/uuid"
)

type CreatePublication struct {
	MaterialType string    `json:"material_type"`
	Status       string    `json:"status"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type Publication struct {
	ID           uuid.UUID
	MaterialType string    `json:"material_type"`
	Status       string    `json:"status"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

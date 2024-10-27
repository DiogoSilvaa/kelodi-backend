package data

import (
	"elodi-backend/internal/validator"
	"time"
)

type Property struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
}

func ValidateProperty(v *validator.Validator, property *Property) {
	v.Check(property.Title != "", "title", "must be provided")
	v.Check(len(property.Title) <= 500, "title", "must not be longer than 500 bytes")

	v.Check(property.Description != "", "description", "must be provided")
	v.Check(len(property.Description) <= 500, "description", "must not be longer than 500 bytes")

	v.Check(property.Location != "", "location", "must be provided")
	v.Check(len(property.Location) <= 500, "location", "must not be longer than 500 bytes")
}

package data

import (
	"database/sql"
	"elodi-backend/internal/validator"
	"errors"
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

type PropertyRepo struct {
	DB *sql.DB
}

func (p PropertyRepo) Insert(property *Property) error {
	query := `
		INSERT INTO properties (title, description, location, created_by)
		VALUES ($1, $2, $3, 'anonymous')
		RETURNING id, created_at, created_by
	`

	args := []interface{}{property.Title, property.Description, property.Location}

	return p.DB.QueryRow(query, args...).Scan(&property.ID, &property.CreatedAt, &property.CreatedBy)
}

func (p PropertyRepo) Get(id int64) (*Property, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `SELECT id, title, description, location, created_at, created_by
		FROM properties
		WHERE id = $1
	`

	var property Property

	err := p.DB.QueryRow(query, id).Scan(&property.ID, &property.Title, &property.Description, &property.Location, &property.CreatedAt, &property.CreatedBy)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &property, nil
}

func (p PropertyRepo) Update(property *Property) error {
	query := `
		UPDATE properties
		SET title = $1, description = $2, location = $3
		WHERE id = $4
	`
	args := []interface{}{
		property.Title,
		property.Description,
		property.Location,
		property.ID,
	}

	_, err := p.DB.Exec(query, args...)

	return err
}

func (p PropertyRepo) Delete(id int64) error {

	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM properties
		WHERE id = $1
	`

	result, err := p.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

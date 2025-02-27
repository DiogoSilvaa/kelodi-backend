package data

import (
	"context"
	"database/sql"
	"kelodi-backend/internal/validator"
	"errors"
	"fmt"
	"time"
)

type Property struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	Version     int64     `json:"version"`
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

func (r PropertyRepo) Insert(property *Property) error {
	query := `
		INSERT INTO properties (title, description, location, created_by)
		VALUES ($1, $2, $3, 'anonymous')
		RETURNING id, created_at, created_by, version
	`

	args := []interface{}{property.Title, property.Description, property.Location}

	return r.DB.QueryRow(query, args...).Scan(&property.ID, &property.CreatedAt, &property.CreatedBy, &property.Version)
}

func (r PropertyRepo) Get(id int64) (*Property, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `SELECT id, title, description, location, created_at, created_by, version
		FROM properties
		WHERE id = $1
	`

	var property Property

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&property.ID,
		&property.Title,
		&property.Description,
		&property.Location,
		&property.CreatedAt,
		&property.CreatedBy,
		&property.Version,
	)
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

func (r PropertyRepo) GetAll(title string, description string, location string, filters Filters) ([]*Property, Metadata, error) {
	query := fmt.Sprintf(`
	SELECT count(*) OVER(), id, title, description, location, created_at, created_by, version
	FROM properties
	WHERE(to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '')
	AND (to_tsvector('simple', description) @@ plainto_tsquery('simple', $2) OR $2 = '')
	AND (to_tsvector('simple', location) @@ plainto_tsquery('simple', $3) OR $3 = '')
	ORDER BY %s %s, id ASC
	LIMIT $4 OFFSET $5
	`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{title, description, location, filters.limit(), filters.offset()}

	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	properties := []*Property{}

	for rows.Next() {
		var property Property

		err := rows.Scan(
			&totalRecords,
			&property.ID,
			&property.Title,
			&property.Description,
			&property.Location,
			&property.CreatedAt,
			&property.CreatedBy,
			&property.Version,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		properties = append(properties, &property)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return properties, metadata, nil
}

func (r PropertyRepo) Update(property *Property) error {
	query := `
		UPDATE properties
		SET title = $1, description = $2, location = $3
		WHERE id = $4 AND version = $5
		RETURNING version
	`
	args := []interface{}{
		property.Title,
		property.Description,
		property.Location,
		property.ID,
		property.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := r.DB.QueryRowContext(ctx, query, args...).Scan(&property.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (r PropertyRepo) Delete(id int64) error {

	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM properties
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := r.DB.ExecContext(ctx, query, id)
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

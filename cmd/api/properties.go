package main

import (
	"elodi-backend/internal/data"
	"elodi-backend/internal/validator"
	"fmt"
	"net/http"
	"time"
)

func (app *application) createPropertyHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Location    string `json:"location"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	v.Check(input.Title != "", "title", "must be provided")
	v.Check(len(input.Title) <= 500, "title", "must not be longer than 500 bytes")

	v.Check(input.Description != "", "description", "must be provided")
	v.Check(len(input.Description) <= 500, "description", "must not be longer than 500 bytes")

	v.Check(input.Location != "", "location", "must be provided")
	v.Check(len(input.Location) <= 500, "location", "must not be longer than 500 bytes")

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) getPropertyHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	property := data.Property{
		ID:          id,
		Title:       "Villa",
		Description: "A 6 Bedroom Villa",
		Location:    "Quinta do Lago",
		CreatedAt:   time.Now(),
		CreatedBy:   "Diogo",
	}

	envelopedData := envelope{"property": property}

	err = app.writeJSON(w, http.StatusOK, envelopedData, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

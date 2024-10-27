package main

import (
	"elodi-backend/internal/data"
	"elodi-backend/internal/validator"
	"fmt"
	"net/http"
	"time"
)

func (app *application) createPropertyHandler(w http.ResponseWriter, r *http.Request) {
	// Decode into intermediary struct to prevent id from being passed in
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

	property := &data.Property{
		Title:       input.Title,
		Description: input.Description,
		Location:    input.Location,
	}

	v := validator.New()
	if data.ValidateProperty(v, property); !v.Valid() {
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

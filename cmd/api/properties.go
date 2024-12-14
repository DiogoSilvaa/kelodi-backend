package main

import (
	"elodi-backend/internal/data"
	"elodi-backend/internal/validator"
	"errors"
	"fmt"
	"net/http"
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

	err = app.models.Properties.Insert(property)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("v1/properties/%d", property.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"property": property}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getPropertyHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	property, err := app.models.Properties.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"property": property}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

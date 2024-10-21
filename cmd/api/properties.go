package main

import (
	"elodi-backend/internal/data"
	"fmt"
	"net/http"
	"time"
)

func (app *application) createPropertyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new property")
}

func (app *application) getPropertyHandler (w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w,r)
		return
	}

	property := data.Property{
		ID: id,
		Title: "Villa",
		Description: "A 6 Bedroom Villa",
		Location: "Quinta do Lago",
		CreatedAt: time.Now(),
		CreatedBy: "Diogo",
	}

	err = app.writeJSON(w, http.StatusOK, property, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
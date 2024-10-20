package main

import (
	"fmt"
	"net/http"
)

func (app *application) createPropertyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new property")
}

func (app *application) fetchPropertyHandler (w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w,r)
		return
	}

	fmt.Fprintf(w, "show the details of movie %d\n", id)
}
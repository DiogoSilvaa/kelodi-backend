package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/properties", app.createPropertyHandler)
	router.HandlerFunc(http.MethodGet, "/v1/properties/:id", app.getPropertyHandler)
	router.HandlerFunc(http.MethodGet, "/v1/properties", app.listPropertiesHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/properties/:id", app.updatePropertyHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/properties/:id", app.deletePropertyHandler)
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activate", app.activateUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.rateLimit(router))
}

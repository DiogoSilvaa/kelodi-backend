package main

import (
	"expvar"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.Handler(http.MethodGet, "/v1/metrics", expvar.Handler())
	router.HandlerFunc(http.MethodGet, "/v1/properties/:id", app.requirePermission("properties:read", app.getPropertyHandler))
	router.HandlerFunc(http.MethodGet, "/v1/properties", app.requirePermission("properties:read", app.listPropertiesHandler))
	router.HandlerFunc(http.MethodPost, "/v1/properties", app.requirePermission("properties:write", app.createPropertyHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/properties/:id", app.requirePermission("properties:write", app.updatePropertyHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/properties/:id", app.requirePermission("properties:write", app.deletePropertyHandler))
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activate", app.activateUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/password", app.updateUserPasswordHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/password-reset", app.createPasswordResetTokenHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/activation", app.createActivationTokenHandler)

	return app.metrics(app.recoverPanic(app.enableCors(app.rateLimit(app.authenticate(router)))))
}

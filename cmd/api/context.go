package main

import (
	"context"
	"kelodi-backend/internal/data"
	"net/http"
)

type contextKey string

const userContextKey = contextKey("user")

// contextSetUser sets the user in the request context.
func (app *application) contextSetUser(r *http.Request, user *data.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)

	return r.WithContext(ctx)
}

// contextGetUser retrieves the user from the request context.
func (app *application) contextGetUser(r *http.Request) *data.User {
	user, ok := r.Context().Value(userContextKey).(*data.User)
	if !ok {
		panic("missing user value in request context")
	}

	return user
}

package main

import (
	"net/http"
	"test/internal/store"
)

func (app *application) listUsersHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement listUsersHandler
}

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement getUserHandler
}

type CreateUserPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateUserPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	User := &store.User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: payload.Password,
	}

	ctx := r.Context()

	if err := app.store.Users.Create(ctx, User); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Something went wrong while creating User, try again.")
		return
	}

	if err := writeJSON(w, http.StatusCreated, User); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Error writing JSON")
		return
	}
}

func (app *application) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement deleteUserHandler
}

func (app *application) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement updateUserHandler
}

package main

import (
	"net/http"
	"test/internal/store"
)

func (app *application) listPostsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement listPostsHandler
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement getPostHandler
}

type CreatePostPayload struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload
	if err := readJSON(w, r, &payload); err != nil {
		writeJSONError(w, http.StatusBadRequest, "Error parsing JSON")
		return
	}

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		// TODO: Change after auth
		Tags:   payload.Tags,
		UserID: "1",
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Something went wrong while creating post, try again.")
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Error writing JSON")
		return
	}
}

func (app *application) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement deletePostHandler
}

func (app *application) updatePostHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement updatePostHandler
}

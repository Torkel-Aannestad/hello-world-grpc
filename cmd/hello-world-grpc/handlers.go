package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (app *application) listPostsHandler(w http.ResponseWriter, r *http.Request) {
	posts := []Post{
		{ID: "1", Title: "first post", Body: "lorum ipsum"},
		{ID: "2", Title: "second post", Body: "lorum ipsum"},
		{ID: "3", Title: "third post", Body: "lorum ipsum"},
	}

	err := app.writeJSON(w, envelope{"posts": posts}, nil, http.StatusOK)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	idParams := chi.URLParamFromCtx(r.Context(), "id")

	//get post with idParams
	post := Post{
		ID:    idParams,
		Title: "Actual Title",
		Body:  "Some body will come",
	}
	err := app.writeJSON(w, envelope{"post": post}, nil, http.StatusOK)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	err := app.readJSON(r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	//swap for real deal
	newPost := Post{
		ID:    "1",
		Title: input.Title,
		Body:  input.Body,
	}

	err = app.writeJSON(w, envelope{"posts": newPost}, nil, http.StatusCreated)
	if err != nil {
		app.logger.Error("error", "msg", err.Error())
		return
	}
}

func (app *application) updatePostHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("updatePostHandler")
	_, err := w.Write([]byte("updatePostHandler response\n"))
	if err != nil {
		app.logger.Error("error", "msg", err.Error())
		return
	}
}

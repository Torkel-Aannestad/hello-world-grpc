package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/posts", app.listPostsHandler)
	r.Get("/posts/{id}", app.getPostHandler)
	r.Post("/posts", app.createPostHandler)
	r.Put("/posts/{id}", app.updatePostHandler)

	return r
}

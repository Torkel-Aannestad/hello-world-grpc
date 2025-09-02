package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("GET /posts", app.listPostsHandler)
	r.HandleFunc("GET /posts/{id}", app.getPostHandler)
	r.HandleFunc("POST /posts", app.createPostHandler)
	r.HandleFunc("PUT /posts/{id}", app.updatePostHandler)

	return r
}

package main

import (
	"fmt"
	"net/http"
)

func (app *application) routes() http.Handler {
	fmt.Println("hi from routes")

	r := http.NewServeMux()
	r.HandleFunc("GET /posts", app.listPostsHandler)
	r.HandleFunc("GET /posts/{id}", app.getPostHandler)
	r.HandleFunc("POST /posts", app.createPostHandler)
	r.HandleFunc("PUT /posts/{id}", app.updatePostHandler)

	return r
}

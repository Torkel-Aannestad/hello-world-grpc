package main

import (
	"fmt"
	"net/http"
)

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (app *application) getPostsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getPostsHandler")
	w.Write([]byte("getPostsHandler handler works"))
}

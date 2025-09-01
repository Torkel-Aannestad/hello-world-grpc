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

func (app *application) getPostsHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("getPostsHandler")
	_, err := w.Write([]byte("getPostsHandler handler works"))
	if err != nil {
		app.logger.Error("error", "msg", err.Error())
		return
	}
}

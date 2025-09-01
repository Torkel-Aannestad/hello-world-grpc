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

func (app *application) listPostsHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("getPostsHandler")
	_, err := w.Write([]byte("getPostsHandler handler works"))
	if err != nil {
		app.logger.Error("error", "msg", err.Error())
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("getPostHandler")
	_, err := w.Write([]byte("getPostHandler response /n"))
	if err != nil {
		app.logger.Error("error", "msg", err.Error())
		return
	}
}

func (app *application) createPostHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("createPostHandler")
	_, err := w.Write([]byte("createPostHandler response /n"))
	if err != nil {
		app.logger.Error("error", "msg", err.Error())
		return
	}
}

func (app *application) updatePostHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("updatePostHandler")
	_, err := w.Write([]byte("updatePostHandler response /n"))
	if err != nil {
		app.logger.Error("error", "msg", err.Error())
		return
	}
}

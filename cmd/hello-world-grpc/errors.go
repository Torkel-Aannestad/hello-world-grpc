package main

import "net/http"

func (app *application) logError(r *http.Request, err error) {
	method := r.Method
	uri := r.URL.RequestURI()

	app.logger.Error(err.Error(), "method", method, "uri", uri)
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, statusCode int, message any) {
	err := app.writeJSON(w, envelope{"error": message}, nil, statusCode)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

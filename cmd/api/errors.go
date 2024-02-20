package main

import (
	"fmt"
	"net/http"
)

func (app *application) LogError(r *http.Request, err error) {
	app.logger.Println(err)
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	envelope := envelope{
		"error": message,
	}

	err := app.writeJSON(w, status, envelope, nil)
	if err != nil {
		app.LogError(r, err)
		w.WriteHeader(500)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	message := "the server encountered a problem and could not process your request."

	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"

	app.errorResponse(w, r, http.StatusNotFound, message)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("method: %s is not supported for this resource.", r.Method)

	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
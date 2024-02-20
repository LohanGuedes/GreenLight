package main

import (
	"net/http"
)

type healthCheckResponse struct {
	Enviroment string `json:"environment"`
	Version    string `json:"version"`
}

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := healthCheckResponse{
		Enviroment: app.config.env,
		Version:    version,
	}

	err := app.writeJSON(w, http.StatusOK, envelope{
		"status":      "avaliable",
		"system_info": data,
	}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

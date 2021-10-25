package main

import (
	"net/http"
)

type healthcheckData struct {
	status      string
	environment string
	version     string
}

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := healthcheckData{
		status:      "available",
		environment: app.config.env,
		version:     version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server ecnountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}
}

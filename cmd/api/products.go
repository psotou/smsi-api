package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/psotou/smsi/internal/data"
)

func (app *application) createProductHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new product")
}

func (app *application) showProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.realIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// New instance of the Product struct
	product := data.Product{
		ID:          id,
		Name:        "MJ230SS",
		Description: "Boquilla misting example",
		CreatedAt:   time.Now(),
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"product": product}, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

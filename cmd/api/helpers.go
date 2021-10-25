package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Retrieve the "id" URL parameter from the current request context, then convert it
// to an integer. If no such parameter exists, or if the parameter is not an integer,
// return 0 and an error.
func (app *application) realIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

type envelope map[string]interface{}

// writeJSON() helper for sending responses that takes the destination
// http.ResponseWriter, the HTTP status code to send, the data to encode to JSON,
// and a header map containing any additional HTTP headers we want to include in the response
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	jsn, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Append a newline to make it easier to view in a terminal application
	jsn = append(jsn, '\n')

	// We loop through the header map and add each header to the http.ResponseWriter header map.
	// Note: Go doesn't through any error when trying to range over (or generally read from) a nil map.
	for key, value := range headers {
		w.Header()[key] = value
	}

	// Add the "Content-Type: application/json" header, then write the status code and JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsn)

	return nil
}

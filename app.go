package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// Application defines the core of the web service.
type Application struct {
	token string
	db    *Database
}

// Response represents the structure of a JSON response.
type Response struct {
	Ok    bool        `json:"ok"`
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

// ResponseAvailability represents the HTTP response for the GET request.
type ResponseAvailability struct {
	Slots []string `json:"slots"`
}

// ResponseAppointment represents the HTTP response for the POST request.
type ResponseAppointment struct {
	UUID string `json:"appointmentId"`
}

// NewApp initializes the whole RESTful web API.
func NewApp(token string) *Application {
	app := new(Application)
	app.db = NewDatabase()
	app.token = token
	return app
}

// Write writes a JSON-encoded object back to the http client.
func (app *Application) Write(w http.ResponseWriter, r *http.Request, v interface{}) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("E_JSON_WRITE", err)
	}
}

// Fail returns a "400 Bad Request" HTTP response with some JSON information.
func (app *Application) Fail(w http.ResponseWriter, r *http.Request, err error) {
	v := Response{Error: err.Error()}

	w.WriteHeader(http.StatusBadRequest)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("E_JSON_WRITE", err)
	}
}

// CheckToken compares the token from the HTTP request with the database.
func (app *Application) CheckToken(token string) error {
	if token == "" {
		return errors.New("token is missing")
	}

	if token == app.token {
		return nil
	}

	return errors.New("token is invalid")
}

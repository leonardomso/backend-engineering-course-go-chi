package main

import (
	"fmt"
	"net/http"
)

// Handles health check requests
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

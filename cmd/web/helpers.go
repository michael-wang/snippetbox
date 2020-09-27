package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// Common code to handle server internal error:
// 1. Print trace.
// 2. Response with what's wrong to client.
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// Common code for client side error:
// 1. Response standard error message of the status.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// For consistency (with http.NotFound), add notFound handler.
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

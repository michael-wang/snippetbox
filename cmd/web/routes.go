package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	// Create a file server to serve files out of "./ui/static" directory.
	// Note: the path to http.Dir is relative to project root.
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// To register the file server as the handler for all URL paths
	// that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}

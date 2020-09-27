package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// application struct holds application-wide dependencies.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// addr contains command-line flag 'addr'.
	addr := flag.String("addr", ":4000", "Address for server to listen to")

	// Must call Parse after setup all flags, and before using them.
	flag.Parse()

	// With custom logs redirect to standard outputs we can:
	// `go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log`
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

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

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// application struct holds application-wide dependencies.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// Parsing runtime configuration settings.
	// addr contains command-line flag 'addr'.
	addr := flag.String("addr", ":4000", "Address for server to listen to")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	// Must call Parse after setup all flags, and before using them.
	flag.Parse()

	// Establishing dependencies for HTTP handlers. (loggers for now.)
	// With custom logs redirect to standard outputs we can:
	// `go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log`
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Running the HTTP server.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}

package main

import (
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)
	svr := newTestServer(t, app.routes())
	defer svr.Close()

	code, _, body := svr.get(t, "/ping")

	if code != http.StatusOK {
		t.Errorf("status code want: %d; got: %d", http.StatusOK, code)
	}

	if string(body) != "OK" {
		t.Errorf("body want: %s; got: %s", "OK", body)
	}
}

package main

import (
	"bytes"
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

func TestShowSnippet(t *testing.T) {
	app := newTestApplication(t)

	svr := newTestServer(t, app.routes())
	defer svr.Close()

	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody []byte
	}{
		{"Valid ID", "/snippet/1", http.StatusOK, []byte("And old slient pond...")},
		{"Non-existed ID", "/snippet/2", http.StatusNotFound, nil},
		{"Negative ID", "/snippet/-1", http.StatusNotFound, nil},
		{"Decimal ID", "/snippet/1.23", http.StatusNotFound, nil},
		{"String ID", "/snippet/foo", http.StatusNotFound, nil},
		{"Empty ID", "/snippet/", http.StatusNotFound, nil},
		{"Trailing splash", "/snippet/1/", http.StatusNotFound, nil},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			code, _, body := svr.get(t, c.urlPath)

			if code != c.wantCode {
				t.Errorf("resp code want: %d; got: %d", c.wantCode, code)
			}

			if !bytes.Contains(body, c.wantBody) {
				t.Errorf("body want to contain: %s but not found", c.wantBody)
			}
		})
	}
}

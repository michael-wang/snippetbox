package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecureHeaders(t *testing.T) {
	w := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(ping)

	secureHeaders(next).ServeHTTP(w, r)
	resp := w.Result()
	defer resp.Body.Close()

	frameOptions := resp.Header.Get("X-Frame-Options")
	if frameOptions != "deny" {
		t.Errorf("frameOptions want %q; got: %q", "deny", frameOptions)
	}

	xssProtection := resp.Header.Get("X-XSS-Protection")
	if xssProtection != "1; mode=block" {
		t.Errorf("xssProtection want: %q; got: %q", "1; mode=block", xssProtection)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status code want: %d; got: %d", http.StatusOK, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "OK" {
		t.Errorf("body want: %s; got: %s", "OK", body)
	}
}

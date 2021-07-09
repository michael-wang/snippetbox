package main

import (
	"bytes"
	"net/http"
	"net/url"
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

func TestSignupUser(t *testing.T) {
	app := newTestApplication(t)
	svr := newTestServer(t, app.routes())
	defer svr.Close()

	_, _, body := svr.get(t, "/user/signup")
	csrfToken := extractCSRFToken(t, body)

	tests := []struct {
		name         string
		userName     string
		userEmail    string
		userPassword string
		csrfToken    string
		wantCode     int
		wantBody     []byte
	}{
		{"Valid submission", "Bob", "bob@example.com", "validPa$$word", csrfToken, http.StatusSeeOther, nil},
		{"Empty name", "", "bob@example.com", "validPa$$word", csrfToken, http.StatusOK, []byte("cannot be blank")},
		{"Empty email", "Bob", "", "validPa$$word", csrfToken, http.StatusOK, []byte("cannot be blank")},
		{"Empty password", "Bob", "bob@example.com", "", csrfToken, http.StatusOK, []byte("cannot be blank")},
		// the regexp allows to match 'bob@example'.
		//{"Invalid email (incomplete domain)", "Bob", "bob@example", "validPa$$word", csrfToken, http.StatusOK, []byte("This field is invalid")},
		{"Invalid email (missing @)", "Bob", "bobexample.com", "validPa$$word", csrfToken, http.StatusOK, []byte("This field is invalid")},
		{"Invalid email (incomplete local part)", "Bob", "@example", "validPa$$word", csrfToken, http.StatusOK, []byte("This field is invalid")},
		{"Short password", "Bob", "bob@example.com", "pa$$word", csrfToken, http.StatusOK, []byte("too short (min 10 characters)")},
		{"Duplicate email", "Bob", "dupe@example.com", "validPa$$word", csrfToken, http.StatusOK, []byte("Address is already in use")},
		{"Invalid CSRF Token", "", "", "", "wrongToken", http.StatusBadRequest, nil},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			form := url.Values{}
			form.Add("name", c.userName)
			form.Add("email", c.userEmail)
			form.Add("password", c.userPassword)
			form.Add("csrf_token", c.csrfToken)

			code, _, body := svr.postForm(t, "/user/signup", form)

			if code != c.wantCode {
				t.Errorf("code want: %d; got: %d", c.wantCode, code)
			}

			if !bytes.Contains(body, c.wantBody) {
				t.Errorf("body want contains %q; but got: %s", c.wantBody, body)
			}
		})
	}
}

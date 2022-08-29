package main

import (
	"bytes"
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {

	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	if code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, code)
	}

	if string(body) != "OK" {
		t.Errorf("want body to equal %q", "OK")
	}
}

func TestShowSnippet(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody []byte
	}{
		{"Valid ID", "/sec/1", http.StatusOK, []byte("An old silent pond...")},
		{"Non-existent ID", "/sec/2", http.StatusNotFound, nil},
		{"Negative ID", "/sec/-1", http.StatusNotFound, nil},
		{"Decimal ID", "/sec/1.23", http.StatusNotFound, nil},
		{"String ID", "/sec/foo", http.StatusNotFound, nil},
		{"Empty ID", "/sec/", http.StatusNotFound, nil},
		{"Trailing slash", "/sec/1/", http.StatusNotFound, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.urlPath)

			if code != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, code)

			}
			if !bytes.Contains(body, tt.wantBody) {
				t.Errorf("want body to contain %q", tt.wantBody)
			}
		})
	}
}

func TestSignupUser(t *testing.T) {
	// Create the application struct containing our mocked dependencies and set
	// up the test server for running and end-to-end test.
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()
	// Make a GET /user/signup request and then extract the CSRF token from the
	// response body.
	_, _, body := ts.get(t, "/user/signup")
	csrfToken := extractCSRFToken(t, body)
	// Log the CSRF token value in our test output. To see the output from the
	// t.Log() command you need to run `go test` with the -v (verbose) flag
	// enabled.
	t.Log(csrfToken)
}

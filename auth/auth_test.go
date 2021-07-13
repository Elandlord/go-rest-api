package auth

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestAuthenticateWithoutCredentials(t *testing.T) {
	writer := httptest.NewRecorder()

	request := httptest.NewRequest("POST", "/authenticate", nil)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	Authenticate(writer, request)

	if writer.Code != http.StatusBadRequest {
		t.Fatalf("Something went wrong. Should receive http.StatusBadRequest.")
	}
}

func TestAuthenticateWithIncorrectCredentials(t *testing.T) {
	writer := httptest.NewRecorder()

	form := url.Values{}
	form.Add("application_name", "test")
	form.Add("password", "test")

	request := httptest.NewRequest("POST", "/authenticate", strings.NewReader(form.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	Authenticate(writer, request)

	if writer.Code != http.StatusUnauthorized {
		t.Fatalf("Something went wrong. Should receive http.StatusUnauthorized.")
	}
}

func TestAuthenticateWithCredentials(t *testing.T) {
	writer := httptest.NewRecorder()

	form := url.Values{}
	form.Add("application_name", "eric")
	form.Add("password", "test")

	request := httptest.NewRequest("POST", "/authenticate", strings.NewReader(form.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	Authenticate(writer, request)

	if writer.Code != http.StatusOK {
		t.Fatalf("Something went wrong. Should receive http.StatusOK.")
	}
}

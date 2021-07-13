package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"mentechmedia.nl/rest-api/config"
)

func TestAllArticlesWithoutToken(t *testing.T) {
	configFile := config.GetTestingConfig()
	db := config.DbConnect(configFile)

	writer := httptest.NewRecorder()

	request := httptest.NewRequest("GET", "/articles", nil)

	AllArticles(db, writer, request)

	if writer.Code != http.StatusOK {
		t.Fatalf("Something went wrong. Should receive http.StatusOK.")
	}
}

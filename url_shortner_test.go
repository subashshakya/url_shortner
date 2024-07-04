package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

var url_code map[string]interface{}

func TestShortenURLHandler(t *testing.T) {
	jsonStr := []byte(`{"url": "https://project-one-two-three-to-the-four"}`)

	req, err := http.NewRequest("POST", "http://localhost:3000/shorten", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(NewRouter().ServeHTTP)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expectedPattern := `http://localhost:3000/[a-zA-Z0-9]{8}`
	error := json.NewDecoder((rr.Body)).Decode(&url_code)
	if error != nil {
		log.Print("Error decoding json body", error)
		return
	}
	log.Print(url_code["short_url"])
	match, err := regexp.MatchString(expectedPattern, string(url_code["short_url"].(string)))
	if err != nil {
		t.Fatal(err)
	}
	if !match {
		t.Errorf("handler returned unexpected body format: got %v, want match to pattern %v",
			rr.Body.String(), expectedPattern)
	}
}

func TestRedirectHandler(t *testing.T) {
	req, err := http.NewRequest("GET", url_code["short_url"].(string), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(NewRouter().ServeHTTP)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusFound)
	}
}

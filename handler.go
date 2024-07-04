package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if !isValidURL(req.URL) {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	shortCode := generateShortCode()
	storage[shortCode] = req.URL

	resp := ShortenResponse{ShortURL: "http://localhost:3000/" + shortCode}
	respondWithJSON(w, http.StatusOK, resp)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortCode := strings.TrimPrefix(r.URL.Path, "/")
	if shortCode == "" || shortCode == "shorten" {
		http.NotFound(w, r)
		return
	}
	if originalURL, found := storage[shortCode]; found {
		http.Redirect(w, r, originalURL, http.StatusFound)
		return
	}
	http.NotFound(w, r)
}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

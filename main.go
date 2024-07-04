package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	log.Println("Starting server on :3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}

package main

import (
	"math/rand"
	"net/url"
	"time"
)

const (
	shortCodeLength = 8
	base62Chars     = "9fYgiSHqMdKx0tZuvjd0Dipm6dDakK9xPh7DYBzpsfRUFaRtjkmg0hldBbX9nH"
)

func isValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}

func generateShortCode() string {
	rand.Seed(time.Now().UnixNano())
	code := make([]byte, shortCodeLength)
	for i := range code {
		code[i] = base62Chars[rand.Intn(len(base62Chars))]
	}
	return string(code)
}

package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"sync"
)

type URLShortener struct {
    urls  map[string]string
    mutex sync.RWMutex
}

func NewURLShortener() *URLShortener {
    return &URLShortener{
        urls: make(map[string]string),
    }
}

// Generate a random short code
func generateShortCode() (string, error) {
    b := make([]byte, 6)
    _, err := rand.Read(b)
    if err != nil {
        return "", err
    }
    return base64.URLEncoding.EncodeToString(b)[:6], nil
}

// Handler for creating short URLs
func (u *URLShortener) shortenHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    longURL := r.FormValue("url")
    if longURL == "" {
        http.Error(w, "URL is required", http.StatusBadRequest)
        return
    }

    shortCode, err := generateShortCode()
    if err != nil {
        http.Error(w, "Error generating short code", http.StatusInternalServerError)
        return
    }

    u.mutex.Lock()
    u.urls[shortCode] = longURL
    u.mutex.Unlock()

    fmt.Fprintf(w, "Short URL: http://localhost:8080/%s\n", shortCode)
}

// Handler for redirecting short URLs
func (u *URLShortener) redirectHandler(w http.ResponseWriter, r *http.Request) {
    shortCode := r.URL.Path[1:]
    
    u.mutex.RLock()
    longURL, exists := u.urls[shortCode]
    u.mutex.RUnlock()

    if !exists {
        http.Error(w, "URL not found", http.StatusNotFound)
        return
    }

    http.Redirect(w, r, longURL, http.StatusMovedPermanently)
}

// Main function to start the HTTP server
func main() {
	urlShortener := NewURLShortener()

	http.HandleFunc("/shorten", urlShortener.shortenHandler)  // Endpoint to shorten URL
	http.HandleFunc("/", urlShortener.redirectHandler)       // Endpoint to redirect short URL

	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
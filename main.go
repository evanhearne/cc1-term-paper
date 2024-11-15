package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// URLStore stores the original URLs mapped to shortened versions
type URLStore struct {
	mapping map[string]string
	mutex   sync.RWMutex
}

// NewURLStore creates and returns a new URLStore
func NewURLStore() *URLStore {
	return &URLStore{
		mapping: make(map[string]string),
	}
}

// GenerateShortURL generates a random string to use as the shortened URL
func GenerateShortURL() string {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal("error generating short URL:", err)
	}
	return base64.URLEncoding.EncodeToString(b)
}

// ShortenURL adds the original URL to the store and returns a shortened URL
func (store *URLStore) ShortenURL(originalURL string) string {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	// Generate and store the shortened URL
	shortURL := GenerateShortURL()
	store.mapping[shortURL] = originalURL
	return shortURL
}

// RetrieveURL fetches the original URL from the store using the shortened URL
func (store *URLStore) RetrieveURL(shortURL string) (string, bool) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	originalURL, exists := store.mapping[shortURL]
	return originalURL, exists
}

// ShortenHandler handles the creation of shortened URLs
func (store *URLStore) ShortenHandler(w http.ResponseWriter, r *http.Request) {
	originalURL := r.URL.Query().Get("url")
	if originalURL == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}
	shortURL := store.ShortenURL(originalURL)
	fmt.Fprintf(w, "Shortened URL: http://localhost:8080/%s\n", shortURL)
}

// RedirectHandler handles the redirection from short URL to the original URL
func (store *URLStore) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:] // Get the URL path without the leading "/"
	originalURL, exists := store.RetrieveURL(shortURL)
	if !exists {
		http.Error(w, "Shortened URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, originalURL, http.StatusFound)
}

func main() {
	store := NewURLStore()

	http.HandleFunc("/shorten", store.ShortenHandler)
	http.HandleFunc("/", store.RedirectHandler)

	fmt.Println("URL shortener running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

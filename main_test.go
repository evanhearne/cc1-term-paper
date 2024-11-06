package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// TestShortenHandler tests that ShortenHandler returns a shortened URL when given an original URL.
func TestShortenHandler(t *testing.T) {
	store := NewURLStore()
	server := httptest.NewServer(http.HandlerFunc(store.ShortenHandler))
	defer server.Close()

	originalURL := "http://example.com"
	req, err := http.NewRequest("GET", server.URL+"/shorten?url="+url.QueryEscape(originalURL), nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", resp.Status)
	}

	// Read the response body correctly
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	body := string(bodyBytes)

	// Verify response contains the "Shortened URL" string
	if !strings.Contains(body, "Shortened URL") {
		t.Errorf("Expected response to contain 'Shortened URL'; got %v", body)
	}
}

// TestRedirectHandler tests that RedirectHandler correctly redirects to the original URL.
func TestRedirectHandler(t *testing.T) {
	store := NewURLStore()

	// Shorten a URL
	originalURL := "http://example.com"
	shortURL := store.ShortenURL(originalURL)

	// Create request to the shortened URL path
	req := httptest.NewRequest("GET", "/"+shortURL, nil)
	w := httptest.NewRecorder()

	// Invoke the RedirectHandler
	store.RedirectHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusFound {
		t.Errorf("Expected status 302 Found; got %v", resp.StatusCode)
	}

	// Check if the Location header matches the original URL
	if location := resp.Header.Get("Location"); location != originalURL {
		t.Errorf("Expected redirect location to be %v; got %v", originalURL, location)
	}
}

// TestRetrieveURLNotFound tests that RedirectHandler returns a 404 when the shortened URL does not exist.
func TestRetrieveURLNotFound(t *testing.T) {
	store := NewURLStore()

	req := httptest.NewRequest("GET", "/nonexistent", nil)
	w := httptest.NewRecorder()

	store.RedirectHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		t.Errorf("Expected status 404 Not Found; got %v", resp.StatusCode)
	}
}

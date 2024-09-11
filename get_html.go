package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("couldn't fetch the URL %s: %w", rawURL, err)
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return "", fmt.Errorf("error status code %d for URL %s", res.StatusCode, rawURL)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("expected content type text/html, got %s", contentType)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't read HTML body from URL %s: %w", rawURL, err)
	}

	return string(body), nil
}

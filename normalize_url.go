package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(initialURL string) (string, error) {
	u, err := url.Parse(strings.TrimRight(initialURL, "/"))
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %v", err)
	}

	resultURL := u.Hostname() + u.Path
	resultURL = strings.TrimPrefix(resultURL, "www.")
	resultURL = strings.ToLower(resultURL)
	return resultURL, nil
}

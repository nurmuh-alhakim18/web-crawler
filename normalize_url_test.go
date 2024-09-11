package main

import (
	"strings"
	"testing"
)

func TestNormalizeUrl(t *testing.T) {
	tests := []struct {
		name       string
		initialURL string
		resultURL  string
	}{
		{
			name:       "remove scheme",
			initialURL: "https://leagueoflegends.com/en-us/champions",
			resultURL:  "leagueoflegends.com/en-us/champions",
		},
		{
			name:       "remove extra '/'",
			initialURL: "https://leagueoflegends.com/en-us/champions/",
			resultURL:  "leagueoflegends.com/en-us/champions",
		},
		{
			name:       "lower capital URL",
			initialURL: "https://LEAGUEOFLEGENDS.com/en-us/champions",
			resultURL:  "leagueoflegends.com/en-us/champions",
		},
		{
			name:       "remove scheme and capitals and trailing slash",
			initialURL: "https://LEAGUEOFLEGENDS.com/en-us/champions/",
			resultURL:  "leagueoflegends.com/en-us/champions",
		},
		{
			name:       "remove www",
			initialURL: "https://www.LEAGUEOFLEGENDS.com/en-us/champions/",
			resultURL:  "leagueoflegends.com/en-us/champions",
		},
		{
			name:       "handle invalid URL",
			initialURL: `:\\invalidURL`,
			resultURL:  "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.initialURL)
			if err != nil && !strings.Contains(err.Error(), "couldn't parse URL") {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
			}

			if actual != tc.resultURL {
				t.Errorf("Test %v - '%s' FAIL: expected URL: %v, actual: %v", i, tc.name, tc.resultURL, actual)
			}
		})
	}
}

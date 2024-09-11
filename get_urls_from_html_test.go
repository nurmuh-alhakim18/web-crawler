package main

import (
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name        string
		inputURL    string
		inputBody   string
		expected    []string
		expectError bool
	}{
		{
			name:     "absolute URL",
			inputURL: "https://leagueoflegends.com",
			inputBody: `
<html>
	<body>
		<a href="https://leagueoflegends.com">
			<span>LoL</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://leagueoflegends.com"},
		},
		{
			name:     "relative URL",
			inputURL: "https://leagueoflegends.com",
			inputBody: `
<html>
	<body>
		<a href="/en-us/champions">
			<span>LoL Champions</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://leagueoflegends.com/en-us/champions"},
		},
		{
			name:     "absolute and relative URLs",
			inputURL: "https://leagueoflegends.com",
			inputBody: `
<html>
	<body>
		<a href="/en-us/champions">
			<span>LoL Champions</span>
		</a>
		<a href="https://leagueoflegends.com">
			<span>LoL</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://leagueoflegends.com/en-us/champions", "https://leagueoflegends.com"},
		},
		{
			name:     "no href",
			inputURL: "https://leagueoflegends.com",
			inputBody: `
<html>
	<body>
		<a>
			<span>LoL></span>
		</a>
	</body>
</html>
`,
			expected: nil,
		},
		{
			name:     "bad HTML",
			inputURL: "https://leagueoflegends.com",
			inputBody: `
<html body>
	<a href="en-us/champions">
		<span>LoL Champions></span>
	</a>
</html body>
`,
			expected: []string{"https://leagueoflegends.com/en-us/champions"},
		},
		{
			name:     "invalid href URL",
			inputURL: "https://leagueoflegends.com",
			inputBody: `
<html>
	<body>
		<a href=":\\invalidURL">
			<span>LoL</span>
		</a>
	</body>
</html>
`,
			expected: nil,
		},
		{
			name:     "handle invalid base URL",
			inputURL: `:\\invalidBaseURL`,
			inputBody: `
<html>
	<body>
		<a href="/path">
			<span>LoL</span>
		</a>
	</body>
</html>
`,
			expected:    nil,
			expectError: true,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				if !tc.expectError {
					t.Errorf("Test %v - '%s' FAIL: couldn't parse input URL: %v", i, tc.name, err)
				}
				return
			}

			actual, err := getURLsFromHTML(tc.inputBody, baseURL)
			if err != nil && !strings.Contains(err.Error(), "couldn't parse base URL") {
				t.Errorf("error: %v", err)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - '%s' FAIL: expected URLs: %v, actual: %v", i, tc.name, tc.expected, actual)
				return
			}
		})
	}
}

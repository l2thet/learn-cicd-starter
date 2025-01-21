package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		name     string
		headers  map[string][]string
		expected string
		err      error
	}

	tests := []test{
		{
			name:     "no auth header",
			headers:  map[string][]string{},
			expected: "",
			err:      ErrNoAuthHeaderIncluded,
		},
		{
			name:     "malformed auth header",
			headers:  map[string][]string{"Authorization": {"Bearer"}},
			expected: "",
			err:      errors.New("malformed authorization header"),
		},
		{
			name:     "valid auth header",
			headers:  map[string][]string{"Authorization": {"ApiKey "}},
			expected: "",
			err:      nil,
		}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			headers := http.Header(tc.headers)
			apiKey, err := GetAPIKey(headers)
			if apiKey != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, apiKey)
			}
			if err != nil && err.Error() != tc.err.Error() {
				t.Errorf("expected %v, got %v", tc.err, err)
			}
		})
	}
}

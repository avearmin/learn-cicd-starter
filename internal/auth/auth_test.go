package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input http.Header
		want  string
	}{
		"simple": {
			input: http.Header{
				"Authorization": []string{"ApiKey DUMMY_API_KEY_1234567890"},
			},
			want: "DUMMY_API_KEY_1234567890",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := GetAPIKey(test.input)
			if test.want != got {
				t.Fatalf("expected: %[1]T | %[1]v, got: %[2]T | %[2]v", test.want, got)
			}
		})
	}
}

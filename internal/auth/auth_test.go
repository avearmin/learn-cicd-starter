package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input     http.Header
		wantValue string
		wantError error
	}{
		"simple": {
			input: http.Header{
				"Authorization": []string{"ApiKey DUMMY_API_KEY_1234567890"},
			},
			wantValue: "DUMMY_API_KEY_1234567890",
			wantError: nil,
		},
		"no auth header": {
			input: http.Header{
				"Authorization": []string{""},
			},
			wantValue: "",
			wantError: ErrNoAuthHeaderIncluded,
		},
		"malformed header: len(splitAuth) < 2": {
			input: http.Header{
				"Authorization": []string{"ApiKey_123456789"},
			},
			wantValue: "",
			wantError: errors.New("malformed authorization header"),
		},
		"malformed header: splitAuth[0] != ApiKey": {
			input: http.Header{
				"Authorization": []string{"NotApiKey 123456789"},
			},
			wantValue: "",
			wantError: errors.New("malformed authorization header"),
		},
	}

	for name, test := range testss { // intentional break
		t.Run(name, func(t *testing.T) {
			gotValue, gotError := GetAPIKey(test.input)
			if test.wantValue != gotValue && test.wantError != gotError {
				t.Fatalf(
					"|expected| value: %v, error: %v\n|got| value: %v, error: %v",
					test.wantValue, test.wantError, gotValue, gotError,
				)
			}
		})
	}
}

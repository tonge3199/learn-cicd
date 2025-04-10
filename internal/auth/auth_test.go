package auth

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		header http.Header
		want   string
		err    error
	}{
		"normal": {header: http.Header{"Connection": {"Keep-Alive"},
			"Content-Type":  {"text/html"},
			"Authorization": {"ApiKey the-real-api-key"}},

			want: "the-real-api-key", err: nil,
		},
		"normal2": {header: http.Header{"Connection": {"Keep-Alive"},
			"Content-Type":  {"text/html"},
			"":              {""},
			"Authorization": {"ApiKey the-real-api-key"}},

			want: "the-real-api-key", err: nil,
		},
		"Auth-Bearer": {header: http.Header{"Connection": {"Keep-Alive"},
			"Authorization": {"Bearer someToken"}},

			want: "", err: errors.New("malformed authorization header"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(test.header)

			if got != test.want {
				fmt.Println(err)
				t.Fatalf("expected: %v, got: %v", test.want, got)
			}

			// error check
			if err != nil {
				if err.Error() != test.err.Error() {
					t.Fatalf("expected error: %v, got: %v", test.err, err)
				}
			}
		})
	}
}

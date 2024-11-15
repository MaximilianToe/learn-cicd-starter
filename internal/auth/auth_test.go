package auth

import (
	"errors"
	"net/http"
	"testing"
)

type Tuple struct {
	str string
	err error
}

func TestAuth(t *testing.T) {
	var header http.Header
	gotStr, gotErr := GetAPIKey(header)
	wantStr, wantErr := "", errors.New("malformed authorization header")
	_, _ = Tuple{wantStr, wantErr}, Tuple{gotStr, gotErr}
	// if !(want== got) {
	// 	t.Fatalf("expected: %v, got: %v", want, got)
	// }
}

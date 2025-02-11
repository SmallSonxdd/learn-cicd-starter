package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headerVar := http.Header{}
	gotStr, gotErr := GetAPIKey(headerVar)
	/*if gotErr != nil {
		if gotErr != ErrNoAuthHeaderIncluded {
			fmt.Println("Wrong error")
			return
		}
	}*/
	wantStr, wantErr := "", ErrNoAuthHeaderIncluded
	if !reflect.DeepEqual(wantStr, gotStr) {
		t.Fatalf("expected: %v, got: %v", wantStr, gotStr)
	}
	if !errors.Is(wantErr, gotErr) {
		t.Fatalf("expected: %v, got: %v", wantErr, gotErr)
	}
}

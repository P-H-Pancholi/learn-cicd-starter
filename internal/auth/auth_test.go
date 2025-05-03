package auth

import (
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestSuccessAPIKey(t *testing.T) {

	reqBody := strings.NewReader("")
	req, err := http.NewRequest("GET", "http://example.com", reqBody)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "ApiKey abcdefghijk")

	want := "abcdefghijk"
	got, err := GetAPIKey(req.Header)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestFailAPIKey(t *testing.T) {

	reqBody := strings.NewReader("")
	req, err := http.NewRequest("GET", "http://example.com", reqBody)
	if err != nil {
		t.Fatal(err)
	}

	_, err = GetAPIKey(req.Header)

	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected error : %v, Found error : %v", ErrNoAuthHeaderIncluded, err)
	}
}

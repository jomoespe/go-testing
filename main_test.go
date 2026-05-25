package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Unit test
func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello world!", "!dlrow olleH"},
		{"1234567890", "0987654321"},
		{"", ""},
	}

	for _, tc := range testcases {
		rev := Reverse(tc.in)

		if rev != tc.want {
			t.Errorf("Reverse: got %s, want: %s", rev, tc.want)
		}
	}
}

// TODO Fuzzy test

// Benchmark test
func BenchmarkReverse(b *testing.B) {
	for b.Loop() {
		Reverse("str")
	}
}

// HTTP Handler: Integration test
func TestReverseHandler(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello world!", "!dlrow olleH"},
		//{"1234567890", "0987654321"},
		//{"", ""},
	}

	for _, tc := range testcases {
		r := httptest.NewRequest(http.MethodGet, "/reverse", nil)
		w := httptest.NewRecorder()

		reverseHandler(w, r)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Errorf("/reverse. Wrong status: expected: %v, got: %v", http.StatusOK, res.StatusCode)
		}
		payload, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("/reverse. Error reading response body: %v", err)
		}
		if string(payload) != tc.want {
			t.Fatalf("/reverse. Invalid response payload. Got: %s, want: %s", string(payload), tc.want)

		}
	}
}

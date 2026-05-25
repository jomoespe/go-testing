package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"unicode/utf8"
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
		rev, _ := Reverse(tc.in)

		if rev != tc.want {
			t.Errorf("Reverse: got %s, want: %s", rev, tc.want)
		}
	}
}

// TODO Fuzzy test
func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, _ := Reverse(orig)
		doubleRev, _ := Reverse(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}

// Benchmark test
func BenchmarkReverse(b *testing.B) {
	for b.Loop() {
		Reverse("str")
	}
}

// HTTP Handler: Integration test
func TestReverseHandler(t *testing.T) {
	testcases := []struct {
		str, want string
	}{
		{"", "!dlrow olleH"},
		{"Hello world!", "!dlrow olleH"},
		{"1234567890", "0987654321"},
	}

	for _, tc := range testcases {
		r := httptest.NewRequest(http.MethodGet, "/reverse?q="+url.QueryEscape(tc.str), nil)
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

package stringutils

import (
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

// Benchmark test

func BenchmarkReverse(b *testing.B) {
	for b.Loop() {
		Reverse("str")
	}
}

// Fuzzy test

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

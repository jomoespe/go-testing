package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"unicode/utf8"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /reverse", reverseHandler)

	slog.Info("Listen on 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error("Error starting HTTP server.", slog.Any("error", err))
		os.Exit(1)
	}
}

func reverseHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	if q == "" {
		q = "Hello world!"
	}
	rev, err := Reverse(q)
	if err != nil {
		fmt.Fprint(w, "") // TODO send error
	}
	fmt.Fprint(w, rev)
}

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	b := []rune(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b), nil
}

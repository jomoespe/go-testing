package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/jomoespe/go-testing/pkg/stringutils"
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
	rev, err := stringutils.Reverse(q)
	if err != nil {
		fmt.Fprint(w, "") // TODO send error
	}
	fmt.Fprint(w, rev)
}

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := os.Getenv("NAME")
		version := os.Getenv("VERSION")

		_, _ = fmt.Fprintf(w, "Name: %s, version: %s", name, version)
	})

	_ = http.ListenAndServe(":8080", nil)
}

package main

import (
	"io"
	"net/http"
)

func firstPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello, I am Rush Zhang</h1>")
}

func main() {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":7000", nil)
} 
package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "<h1>Hello World!</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "About")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server Starting...")
	_ = http.ListenAndServe(":3000", nil)
}

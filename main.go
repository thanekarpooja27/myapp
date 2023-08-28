package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World, Welcome to GoLang world!")
	})
	http.ListenAndServe(":8081", nil)
}

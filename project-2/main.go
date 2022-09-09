package main

import (
	"fmt"
	"html"
	"net/http"
)

func test0(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are on the home page.")
}

func test1(w http.ResponseWriter, r *http.Request) {
	path := html.EscapeString(r.URL.Path)
	fmt.Fprintf(w, "Now you are on: %q", path)
}

func test2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Image page")
}

func main() {
	http.HandleFunc("/", test0)
	http.HandleFunc("/about", test1)
	http.HandleFunc("/images", test2)

	http.ListenAndServe(":8080", nil)
}

// Run on -> http://localhost:8080/

package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler1)
	http.HandleFunc("/Hello", handler2)
	http.ListenAndServe(":8080", nil)
}

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome in Golang World\n")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Arpan Welcome to Golang\n")
}

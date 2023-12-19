// main.go
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Docker 2!")
}

func main() {
	http.HandleFunc("/service2", handler)
	http.ListenAndServe(":8080", nil)
}
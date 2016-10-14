package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("HELLO_TO")
	if len(name) == 0 {
		name = "World"
	}
	greeting := fmt.Sprintf("Hello %s!", name)

	io.WriteString(w, greeting)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}

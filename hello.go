package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("HELLO_TO")
	if len(name) == 0 {
		name = "World"
	}

	wait := os.Getenv("ADD_JITTER")
	if len(wait) > 0 {
		s1 := rand.NewSource(time.Now().UnixNano())
		rand1 := rand.New(s1)

		r := rand1.Intn(100)
		r += 100
		time.Sleep(time.Duration(r) * time.Millisecond)
	}

	greeting := fmt.Sprintf("Hello %s!", name)

	io.WriteString(w, greeting)
}

func main() {
	http.HandleFunc("/", hello)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"

	}

	http.ListenAndServe(":"+port, nil)
}

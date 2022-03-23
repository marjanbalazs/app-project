package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello!"))
	})

	err := http.ListenAndServe(":4850", mux)
	if err != nil {
		log.Fatal(err)
	}
}

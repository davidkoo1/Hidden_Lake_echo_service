package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		req, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Println(r.Method, string(req))
		fmt.Fprintf(w, "echo(%s)", string(req))
	})
	log.Println("Service is Listening...")
	http.ListenAndServe("127.0.0.1:8080", nil)
}

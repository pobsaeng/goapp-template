package main

import (
	"log"
	"net/http"

	handler "github.com/gotpl-app/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", logRequest(handler.HandleIndex))
	mux.HandleFunc("/signin", handler.HandleSignIn)
	mux.HandleFunc("/success", logRequest(handler.HandleSuccess))
	log.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func logRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Incoming request: %s %s\n", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	}
}
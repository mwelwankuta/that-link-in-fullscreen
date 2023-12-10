package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := "3000"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	r := mux.NewRouter()
	r.HandleFunc("/watch/{video_link}", HandlePlayRoute).Methods("GET")
	r.HandleFunc("/health", HandleHealthRoute).Methods("GET")
	http.Handle("/", r)

	fmt.Printf("Player running on :%s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}

package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mwelwankuta/that-link-in-fullscreen/api/routes"
)

func StartServer() {
	r := mux.NewRouter()

	r.HandleFunc("/api/watch", routes.OpenVideoLinkInBrowser).Methods("POST")
	r.HandleFunc("/api/pause", routes.PauseVideo).Methods("POST")

	http.Handle("/", r)

	port := "5050"
	envPort := os.Getenv("PORT")
	if envPort != "" {
		port = envPort
	}

	fmt.Printf("Server is running on :%s...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}

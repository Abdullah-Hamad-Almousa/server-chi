package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.Get("/hello", basicHandler)

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to start the server", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	write, err := w.Write([]byte("Hello Nerds!"))
	if err != nil {
		fmt.Println("Failed to write the response", err)
	}
	log.Println(write)
}

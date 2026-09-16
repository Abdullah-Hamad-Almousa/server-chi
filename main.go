package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(basicHandler),
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

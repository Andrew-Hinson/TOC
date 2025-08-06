package main

import (
	"log"
	"net/http"

	"github.com/Andrew-Hinson/toc/services/auth"
)

func main() {
	log.Println("Starting HTTP Server...")
	server := &http.Server{
		Addr:    ":8080",
		Handler: auth.New(),
	}

	log.Printf("Started HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed")
	}
}

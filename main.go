package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Andrew-Hinson/toc/services/auth"
)

func main() {
	fmt.Println("Starting HTTP Server...")
	server := &http.Server{
		Addr:    ":8080",
		Handler: auth.New(),
	}

	fmt.Printf("Started HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed")
	}
}

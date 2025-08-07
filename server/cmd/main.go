package main

import (
	"log"
	"net/http"

	"github.com/Andrew-Hinson/toc/server/server/services/auth"
	"github.com/Andrew-Hinson/toc/server/server/services/chat"
)

func main() {
	log.Println("Starting toc server...")

	hub := chat.NewHub()
	go hub.Run()

	mux := http.NewServeMux()

	authHandler := auth.New()
	mux.Handle("/auth/", authHandler)

	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		chat.ServeWs(hub, w, r)
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is running"))
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Printf("Started toc server. Listening at %q", server.Addr)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed")
	}
}

package main

import (
	"log"
	"net/http"

	"github.com/Andrew-Hinson/toc/server/server/services/auth"
	"github.com/Andrew-Hinson/toc/server/server/services/chat"
)

func main() {
	log.Println("Starting HTTP Server...")

	// Create the websocket hub and start it in a goroutine
	hub := chat.NewHub()
	go hub.Run()

	// Create a combined mux that handles both auth and websocket routes
	mux := http.NewServeMux()

	// Mount auth routes
	authHandler := auth.New()
	mux.Handle("/auth/", authHandler)

	// Add websocket endpoint
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		chat.ServeWs(hub, w, r)
	})

	// Add a simple health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is running"))
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Printf("Started HTTP Server. Listening at %q", server.Addr)
	log.Println("Auth endpoints: /auth/login, /auth/google/login, /auth/google/callback")
	log.Println("WebSocket endpoint: /ws")
	log.Println("Health check: /health")

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed")
	}
}

package main

import (
	"log"
	"net/http"

	"github.com/Andrew-Hinson/toc/server/server/services/chat"
	"github.com/Andrew-Hinson/toc/server/server/services/routes"
	sqliteStore "github.com/Andrew-Hinson/toc/server/server/services/storage/sqlite"
)

func main() {
	log.Println("Starting toc server...")

	db, err := sqliteStore.Open("./data/toc.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := sqliteStore.Migrate(db); err != nil {
		log.Fatal(err)
	}

	hub := chat.NewHub()
	go hub.Run()

	mux := routes.New(hub)

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

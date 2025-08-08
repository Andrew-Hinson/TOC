package routes

import (
	"net/http"

	"github.com/Andrew-Hinson/toc/server/server/services/auth"

	"github.com/Andrew-Hinson/toc/server/server/services/chat"
)

func New(hub *chat.Hub) *http.ServeMux {
	mux := http.NewServeMux()

	auth.Register(mux)

	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		chat.ServeWs(hub, w, r)
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is running"))
	})

	return mux
}

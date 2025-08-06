package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adriancable/webtransport-go"
)

func main() {
	server := &webtransport.Server{
		ListenAddr: ":4433",
		TLSCert:    webtransport.CertFile{Path: "cert.pem"},
		TLSKey:     webtransport.CertFile{Path: "cert.key"},
	}

	http.HandleFunc("/echo", func(rw http.ResponseWriter, r *http.Request) {
		session := r.Body.(*webtransport.Session)
		session.AcceptSession()
		defer session.CloseSession()

		stream, err := session.AcceptStream()
		if err != nil {
			return
		}

		buf := make([]byte, 1024)
		for {
			n, err := stream.Read(buf)
			if err != nil {
				break
			}
			fmt.Printf("Received: %s\n", buf[:n])
			stream.Write(buf[:n])
		}
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	server.Run(ctx)
}

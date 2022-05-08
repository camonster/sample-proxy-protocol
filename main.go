package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/pires/go-proxyproto"
)

func run() error {
	server := http.Server{
		Addr: ":8080",
	}

	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		return err
	}

	proxyListener := &proxyproto.Listener{
		Listener:          ln,
		ReadHeaderTimeout: 10 * time.Second,
	}
	defer proxyListener.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := fmt.Sprintf("RemoteAddr: %s", r.RemoteAddr)
		log.Println(response)
		w.Write([]byte(response))
	})
	if err := server.Serve(proxyListener); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

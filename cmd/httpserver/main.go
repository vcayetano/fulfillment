package main

import (
	"fmt"
	"github.com/rs/cors"
	"github.com/spf13/cast"
	"github.com/vcayetano/fulfillment/specifications"
	"log"
	"net/http"
	"os"
)

type PackServer struct {
	packer specifications.Packer
}

var (
	defaultPort = "8080"
)

func main() {

	port := os.Getenv("PORT")
	tlsEnabled := cast.ToBool(os.Getenv("TLS_ENABLED"))

	if port == "" {
		port = defaultPort
	}

	packingServer := &PackServer{
		packer: &specifications.DefaultPacker{},
		//packer: specifications.NewPacker(packageSizes),
	}

	server := http.NewServeMux()
	server.HandleFunc("/", packingServer.Home)
	server.HandleFunc("/packs", packingServer.Packs)
	server.HandleFunc("/default-packs", packingServer.DefaultPacks)
	handler := cors.Default().Handler(server)

	if !tlsEnabled {
		if err := http.ListenAndServe(fmt.Sprintf(":%s", port), handler); err != nil {
			log.Fatal(err)
		}
		return
	}

	if err := http.ListenAndServeTLS(fmt.Sprintf(":%s", port), "cert.pem", "key.pem", server); err != nil {
		log.Fatal(err)
	}

}

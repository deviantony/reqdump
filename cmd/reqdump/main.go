package main

import (
	"log"
	"github.com/deviantony/reqdump"
	"github.com/deviantony/reqdump/http"
)

func main() {
	listenAddr := reqdump.DefaultListenAddr + ":" + reqdump.DefaultListenPort
	server := http.NewServer()

	log.Printf("[INFO] - Starting reqdump [addr: %s] [version: %s]", listenAddr, reqdump.Version)
	err := server.Start(listenAddr)
	if err != nil {
		log.Fatalf("Unable to start reqdump: %s\n", err)
	}
}

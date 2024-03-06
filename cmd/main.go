package main

import (
    "flag"
    "log"

    "github.com/CraftyLlamaCoalition/CraftyLlamaNotes/internal/server"
)

var (
    port = flag.Int("port", 50051, "Server Port")
)

func main() {
    flag.Parse()
    gatewayServer := server.NewServer(*port)
    log.Printf("Listening on port %d\n", *port) 
    if err := gatewayServer.ListenAndServe(); err != nil {
        log.Fatalf("Failed to listen on grpc gateway server: %v", err)
    }
}

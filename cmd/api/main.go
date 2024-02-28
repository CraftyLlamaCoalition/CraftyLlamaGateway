package main

import (
    "fmt"
    "flag"
    "log"

    "github.com/CraftyLlamaCoalition/CraftyLlamaGateway/internal/server"
)

var (
    port = flag.Int("port", 50051, "Server Port")
)

func main() {
    flag.Parse()
    gatewayServer := server.NewServer(port)

    if err = gatewayServer.ListenAndServer(); err != nil {
        log.Fatalf("Failed to listen on grpc gateway server: %v", err)
    }
}

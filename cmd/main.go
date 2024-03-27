package main

import (
    "flag"
    "log"
    "net"
    "fmt"

    "google.golang.org/grpc"
    "github.com/CraftyLlamaCoalition/CraftyLlamaNotes/api/grpc/notes"
    pb "github.com/CraftyLlamaCoalition/CraftyLlamaProtoGo"
    
)

var (
    port = flag.Int("port", 50051, "Server Port")
)

func main() {
    flag.Parse()
    log.Printf("Listening on port %d\n", *port) 

    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
        return
    }

    gprcServer := grpc.NewServer()
    //register notes service
    noteServer := &notes.GRPCNoteServer{}
    pb.RegisterCriaNotesServiceServer(gprcServer, noteServer)

    if err := gprcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
        return
    }

}

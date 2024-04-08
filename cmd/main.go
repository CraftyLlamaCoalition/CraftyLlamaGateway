package main

import (
    "log"
    "net"
    "fmt"
    "os"
    "strconv"

    "github.com/joho/godotenv"

    "google.golang.org/grpc"
    "github.com/CraftyLlamaCoalition/CraftyLlamaNotes/api/grpc/notes"
    pb "github.com/CraftyLlamaCoalition/CraftyLlamaProtoGo"
    db "github.com/CraftyLlamaCoalition/CraftyLlamaNotes/internal/criaDB"
    
)


func main() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    port, err := strconv.Atoi(os.Getenv("GRPC_PORT"))
    if err != nil {
        log.Fatalf("Failed to get GPRC PORT number")
        return
    }

    log.Printf("Listening on port %d\n", port) 
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
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

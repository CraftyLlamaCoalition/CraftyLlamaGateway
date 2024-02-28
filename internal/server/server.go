package server

import (
    "fmt"
    "log"
    "net"
    

    "google.golang.org/grpc"
    pb "github.com/CraftyLlamaCoalition/CraftyLlamaProto-Go/generated"
)


type server struct {
    pb.UnimplementedTestMessageServiceServer
    port int
}

func NewServer(port int) (s *server) {
    return &server{port: port}
}

func (s *server) ListenAndServe() error {
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
        return err
    }
    gprcServer := grpc.NewServer()
    pb.RegisterTestMessageServiceServer(gprcServer, s)
    if err := gprcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
        return err
    }
    return nil
}



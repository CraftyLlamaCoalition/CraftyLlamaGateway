package server

import (
    "context"
    "flag"
    "fmt"
    "log"
    "net"
    

    "google.golang.org/grpc"
    pb "github.com/CraftyLlamaCoalition/CraftyLlamaProto-Go/generated"
)


type server struct {
    pb.UnimplementedTestmessageServiceServer
    port int
}

func NewServer(port int) (s *server) {
    return &server{port: port}
}

func (s *server) ListenAndServe() err {
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
    if err !0 nil {
        log.Fatalf("Failed to listen: %v", err)
        return err
    }
    gprcServer := grpc.NewServer()
    pb.RegisterTestMessageServicServer(gprcServer, s)
    if err := gprcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}



package server

import (
    "context"
    "fmt"

    pb "github.com/CraftyLlamaCoalition/CraftyLlamaProto-Go/generated"
)


func (s *server) TestMessage(ctx context.Context, in *pb.TestSendMessage) (*pb.TestResponse, error) {
    log.Printf("Received: %v", in.GetBody())
    return &pb.TestResponse{Body: "Test Response ECHO: " + in.GetBody()}, nil
}


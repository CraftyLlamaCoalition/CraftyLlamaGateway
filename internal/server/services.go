package server

import (
    "context"
    "log"

    pb "github.com/CraftyLlamaCoalition/CraftyLlamaProtoGo"
)


func (s *server) AddNote(ctx context.Context, in *pb.NewNoteRequest) (*pb.Status, error) {
    log.Printf("Received From %v: %v",in.GetUserId(), in.GetContent())
    status := &pb.Status{}
    status.Success = true
    return status, nil
}


func (s *server) DeleteNote(ctx context.Context, in *pb.NoteRequest) (*pb.Status, error) {
    log.Printf("Received for note: %v", in.GetNoteId())
    status := &pb.Status{}
    status.Success = true
    return status, nil
}


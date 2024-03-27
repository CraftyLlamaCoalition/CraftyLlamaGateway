package notes

import (
    "context"
    "log"

    pb "github.com/CraftyLlamaCoalition/CraftyLlamaProtoGo"
)

type GRPCNoteServer struct {
    pb.UnimplementedCriaNotesServiceServer
}


func (s *GRPCNoteServer) AddNote(ctx context.Context, in *pb.NewNoteRequest) (*pb.Status, error) {
    log.Printf("Received From %v: %v",in.GetUserId(), in.GetContent())
    status := &pb.Status{}
    status.Success = true
    return status, nil
}


func (s *GRPCNoteServer) DeleteNote(ctx context.Context, in *pb.NoteRequest) (*pb.Status, error) {
    log.Printf("Received for note: %v", in.GetNoteId())
    status := &pb.Status{}
    status.Success = true
    return status, nil
}

func (s *GRPCNoteServer) GetNote(ctx context.Context, in *pb.NoteRequest) (*pb.NoteContent, error) {

    note := &pb.NoteContent{}
    note.Content = "Not impletmented yet"
    return note, nil
}

func (s *GRPCNoteServer) GetAllNotes(ctx context.Context, in *pb.User) (*pb.MultipleNotes, error) {


    note1 := &pb.NoteContent{}
    note2 := &pb.NoteContent{}
    note1.Content = "Not impletmented yet"
    note2.Content = "Not impletmented yet"
    notes := make([]*pb.NoteContent, 0)
    notes = append(notes,note1)
    notes = append(notes,note2)
    // &pb.MultipleNotes
    return &pb.MultipleNotes{Notes:notes}, nil
}


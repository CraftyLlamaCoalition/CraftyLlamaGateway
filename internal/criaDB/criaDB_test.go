package criaDB

import (
    "testing"
    
    sql "github.com/CraftyLlamaCoalition/CraftlyLlamaNotes/internal/criaDB/sqlcDB"
)

var (
    url = "localhost"  
    dbname = "postgres" 
    user = "postgres"
    pass = "pgtestpassword" 
)


func TestDBConnection(t *testing.T) {
    db, err := Connect(url, dbname, user, pass)
    if err != nil {
        t.Fatalf("Error: %v", err)
    }
    if err = db.Close(); err != nil{
        t.Fatalf("Error: %v", err)
    }
}

func TestAddNote(t *testing.T) {

    db, err := Connect(url, dbname, user, pass)
    defer db.Close()
    if err != nil {
        t.Fatalf("Error: %v", err)
    }
    
    queries := sql.New(db.conn)
    
    newNote, err := queries.CreateNote(db.ctx, sql.CreateNoteParams{
        CreatedBy: "Nickolas larson",
        Content: "New note added through test"
    })
    
    if err != nil {
        t.Fatalf("Error: %v", err)
    }

    t.Printf("Created new note: %d\n", newNote.ID)
}

package notes

import (
    "fmt"
    "testing"
    "github.com/CraftyLlamaCoalition/CraftyLlamaNotes/internal/criaDB"
)

var (
    url = "localhost"  
    dbname = "postgres" 
    user = "postgres"
    pass = "pgtestpassword" 
)
func TestNoteManipulation(t *testing.T) {

    db, err := criaDB.Connect(url, dbname, user, pass)
    defer db.Close()
    if err != nil {
        t.Fatalf(fmt.Sprintf("Error: %v", err))
    }

    createdBy := "Nickolas Larson"
    content := "Testing the add note function" 
    
    newNote, err :=  AddNote(db, createdBy, content)
    if err != nil || newNote.Createdby != createdBy || newNote.Content != content {
        t.Fatalf("Failed to add note")
    }

    getNote, err := GetNote(db, newNote.ID)
    if err != nil || getNote.Createdby != createdBy || getNote.Content != content {
        t.Fatalf("Failed to get note")
    }

    updateLLMNoteText := "New message"
    updatedNote, err := UpdateLLMNote(db, getNote.ID, updateLLMNoteText)
    if err != nil || updatedNote.Createdby != createdBy || updatedNote.Llmresp.String != updateLLMNoteText {
        t.Fatalf(fmt.Sprintf(
            "\nFailed to update note\nNewNote: %d\nGetNote: %d\nUpdatedNote: %d\nError: %v", 
            newNote.ID,
            getNote.ID,
            updatedNote.ID,
            err,
        ))
    }

    if err := DeleteNote(db, updatedNote.ID); err != nil{
        t.Fatalf("Failed to delete note")
    }

} 

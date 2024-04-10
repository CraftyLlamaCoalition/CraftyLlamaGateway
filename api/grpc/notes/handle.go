package notes

import (
    "fmt"
    "log"
    "errors"

    "github.com/CraftyLlamaCoalition/CraftyLlamaNotes/internal/criaDB"
    sql "github.com/CraftyLlamaCoalition/CraftyLlamaNotes/internal/criaDB/sqlcDB"
    "github.com/jackc/pgx/v5/pgtype"
)

func GetNote(db *criaDB.DBConnection, id int32) (*sql.Note, error) {
    queries := sql.New(db.Conn)
    note, err := queries.GetNote(db.Ctx, id)
    if err != nil {
        return nil, errors.New(fmt.Sprintf("Error: %v", err))
    }
    return &note, nil
}

func AddNote(db *criaDB.DBConnection, createdBy string, content string) (*sql.Note, error){

    queries := sql.New(db.Conn)

    newNote, err := queries.CreateNote(db.Ctx, sql.CreateNoteParams{
        Createdby: createdBy,
        Content: content,
    })

    if err != nil {
        return nil, errors.New(fmt.Sprintf("Error: %v", err))
    }

    log.Printf("Created new note: %d\n", newNote.ID)
    return &newNote, nil
}


func UpdateLLMNote(db *criaDB.DBConnection, id int32, llmResp string) (*sql.Note, error) { 

    queries := sql.New(db.Conn)

    data := pgtype.Text{ String: llmResp, Valid: true}

    updatedNote, err := queries.UpdateLLMResp(db.Ctx, sql.UpdateLLMRespParams{
        ID: id,
        Llmresp: data,
    })

    if err != nil {
        return nil, errors.New(fmt.Sprintf("Error: %v", err))
    }

    log.Printf("Created new note: %d\n", updatedNote.ID)
    return &updatedNote, nil
}

func DeleteNote(db *criaDB.DBConnection, id int32) error {
    queries := sql.New(db.Conn)
    return queries.DeleteNote(db.Ctx, id)
}

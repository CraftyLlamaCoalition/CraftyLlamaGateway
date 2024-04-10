package criaDB

import (
    "testing"
    
    // sql "github.com/CraftyLlamaCoalition/CraftyLlamaNotes/internal/criaDB/sqlcDB"
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


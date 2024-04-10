package criaDB

import (
    "fmt"
    "errors"
    "log"
    "context"
    
    "github.com/jackc/pgx/v5"
    // "github.com/jackc/pgx/v5/pgtype"
    //

)

type DBConnection struct {
    Ctx context.Context
    Conn *pgx.Conn
    Name string
}

func Connect(dbURL string, dbname string,  user string, password string) (*DBConnection, error) {
    ctx := context.Background()
    connStr := fmt.Sprintf("postgres://%s:%s@%s/%s", user, password, dbURL, dbname)
    config, err := pgx.ParseConfig(connStr)
    if err != nil {
        return nil, errors.New(fmt.Sprintf("Failed create config"))
    }
    log.Printf("\nHost: %s\nDBName: %s\nUser: %s", config.Host, config.Database, config.User)
    conn, err := pgx.ConnectConfig(ctx, config)
    if err != nil {
        return nil, errors.New(fmt.Sprintf("Failed to connect to DB: %s", connStr))
}

    log.Printf("Connected to DB %s", dbname)
    return &DBConnection{
        Ctx: ctx,
        Conn: conn,
        Name: dbname,
    }, nil
}

func (db *DBConnection) Close() error {
    return db.Conn.Close(db.Ctx)
}






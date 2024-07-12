package util

import (
	"os"
	"errors"
	"context"
	"database/sql"

	_   "embed"

    "survivor/ui"
    "survivor/sqlc"

	_   "github.com/glebarez/go-sqlite"
	rl  "github.com/gen2brain/raylib-go/raylib"
)


const dbName = "desktop_survivor.db"
var db *sql.DB = nil
var query *ui.Queries = nil


func DBExist() bool {
    _, err := os.Stat("desktop_survivor.db")
    if err != nil {
        if errors.Is(err, os.ErrNotExist) {
            return true
        }
        rl.TraceLog(rl.LogFatal, err.Error())
        os.Exit(1)
    }
    return false
}


func DBConnect() (*sql.DB, error) {
    if db != nil {
        return db, nil
    }
    var err error
    db, err = sql.Open("sqlite", dbName)
    return db, err
}


func DBCreate() {
    ctx := context.Background()
    _, err := db.ExecContext(ctx, sqlc.DBSchema)
    if err != nil {
        rl.TraceLog(rl.LogFatal, err.Error())
        os.Exit(1)
    }
}


func DBQuery() *ui.Queries {
    if query != nil {
        return query
    }
    query = ui.New(db)
    return query
}


func DBDelete() error {
    return os.Remove(dbName)
}

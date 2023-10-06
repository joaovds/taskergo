package database

import (
	"database/sql"
	"log"

  _ "github.com/mattn/go-sqlite3"
)

type Connection interface {
  Connect() (*sql.DB, error)
}

func Connect() (*sql.DB, error) {
  db, err := sql.Open("sqlite3", "./database.db")
  if err != nil {
    log.Fatalln("Error on connect to database:\n", err)
    return nil, err
  }

  if err = db.Ping(); err != nil {
    log.Fatalln("Error on ping to database:\n", err)
    return nil, err
  }

  return db, nil
}


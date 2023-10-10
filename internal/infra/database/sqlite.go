package database

import (
	"database/sql"
	"errors"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type Connection interface {
  Connect() (*sql.DB, error)
}

var (
  once sync.Once
  db   *sql.DB
)

func Connect() (*sql.DB, error) {
  once.Do(func () {
    connection, err := sql.Open("sqlite3", "./database.db")
    if err != nil {
      log.Fatalln("Error on connect to database:\n", err)
      return
    }

    if err = connection.Ping(); err != nil {
      log.Panicln("Error on ping to database:\n", err)
      connection.Close()
      connection = nil
      return
    }

    db = connection
  })

  if db == nil {
    return nil, errors.New("Failed to connect to database")
  }

  return db, nil
}


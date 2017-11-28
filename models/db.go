package models

import (
  "fmt"
  "database/sql"
  _ "github.com/lib/pq"
)

var db *sql.DB

func LoadDB(source string) {
  var err error
  db, err = sql.Open("postgres", source)
  // Check connection error
  if err != nil {
    panic(err)
  }
  // Check ping error
  if err = db.Ping(); err != nil {
    panic(err)
  }

  fmt.Println("You connected to your database.")
}


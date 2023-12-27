package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)


const createTable string =
  "CREATE TABLE IF NOT EXISTS Clients (id INTEGER PRIMARY KEY, name TEXT, email TEXT, phone TEXT)" +
  "CREATE TABLE IF NOT EXISTS Rooms (id INTEGER PRIMARY KEY, room_name TEXT, rate REAL)" +
  "CREATE TABLE IF NOT EXISTS Reservations (id INTEGER PRIMARY KEY, client_id INTEGER, room_id INTEGER, checkin_date TEXT, checkout_date TEXT, FOREIGN KEY(client_id) REFERENCES Clients(id), FOREIGN KEY(room_id) REFERENCES Rooms(id))"



type Sql struct {
  db *sql.DB
  tx *sql.Tx
}

func (s *Sql) Init() {
  var err error
  s.db, err = sql.Open("sqlite3", "./database.db")
  if err != nil {
    log.Fatal(err)
  }
}

func (s *Sql) Close() {
  s.db.Close()
}

func (s *Sql) CreateTables() {
  s.db.Exec(createTable)
}

func (s *Sql) ClientInsertion(c Client) {
  
  tx, err := s.db.Begin()
  if err != nil {
    log.Fatal(err)
  }

  stmt, err := tx.Prepare("insert into Clients(name, email, phone) values(?, ?, ?)")
  if err != nil {
    log.Fatal(err)
  }
  defer stmt.Close()

  _, err = stmt.Exec(c.Name, c.Email, c.Phone)
  if err != nil {
    log.Fatal(err)
  }

  tx.Commit()
  c.inDB = true

}

func (s *Sql) ClientPrint() {


  rows, err := s.db.Query("select client_id, name, phone, email from Clients")
  if err != nil {
    log.Print(err)
    return
  }
  defer rows.Close()

  fmt.Printf("------------------------------------------------------\n")
  for rows.Next() { //linked list
    var client_id int
    var name string
    var phone string
    var email string

    err = rows.Scan(&client_id, &name, &phone, &email)
    if err != nil {
      log.Println(err)
      return
    }
    
    fmt.Printf("|%d|%s |%s |%s|\n", client_id, name, phone, email)
  }

  fmt.Printf("----------------------------------------------------\n")



}














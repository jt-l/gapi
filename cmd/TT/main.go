package main

import (
  "fmt"
  "log"
  "os"
  "github.com/james/TT/http"
  "github.com/james/TT/postgres"
)

//All app dependencies are created and injected here
func main() {

  db, err := postgres.Open(os.Getenv("DATABASE_URL")) //open db connection
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  us := &postgres.UserService{DB: db} //inject db into the UserService

  var h http.Handler //declare Handler

  h.UserHandler = http.NewUserHandler() //create instance of user handler

  h.UserHandler.UserService = us //attach UserService to UserHandler

  server := http.NewServer() //create a new server

  server.Handler = &h  //attach handler to server

  err = server.Open() //start the server
  if err != nil {
    fmt.Println(err)
  }

  select{} // block main from terminating to allow our http server to continue to run
}

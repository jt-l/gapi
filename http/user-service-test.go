package http

import (

  "errors"
  "log"
  "testing"

  "github.com/james/tt"
  "github.com/james/tt/http"
  "github.com/james/tt/mock"

)

func TestUserService_CreateUser(t *testing.T) {
  t.Run("OK", testUserService_CreateUser)
}

func testUserService_CreateUser(t *testing.T) {
  s, c := MustOpenServerClient()
  defer s.Close()

}

type Server struct {
  *http.Server

  Handler *Handler
}

//NewServer returns a new instance of Server.
func NewServer() *Server {
  s := &Server {
    Server: http.NewServer()
    Handler: NewHandler()
  }
  s.Server.Handler = s.Handler.Handler

  //use a random port
  s.Addr = ":0"

  return s
}

func MustOpenServerAndClient() (*Server, *http.Client) {
  //Create and open test server.
  s := NewServer()
  if err := s.Open(); err != nil {
    panic(err)
  }

  //Create a client pointing to the server. 
  c := http.NewClient()
  c.URL = url.URL{Scheme: "http", Host: fmt.Sprintf("localhost:%d", s.Port())}

  return s, c
}

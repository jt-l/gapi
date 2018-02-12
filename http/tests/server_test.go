package http_test

import (
  "fmt"
  "net/url"
  "testing"
  "github.com/james/TT/http"
  "io"
  "os"
)

type Server struct {
  *http.Server

  Handler *Handler
}

//NewServer returns a new instance of Server.
func NewServer() *Server {
  s := &Server {
    Server: http.NewServer(),
    Handler: NewHandler(),
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

func VerboseWriter(w io.Writer) io.Writer {
  if testing.Verbose() {
    return io.MultiWriter(w, os.Stderr)
  }
  return w
}

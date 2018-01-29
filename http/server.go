package http

import (
  "net"
  "net/http"
  "os"
  "net/url"
  "github.com/james/TT"
)

//return the default address
var DefaultAddr = ":" + os.Getenv("PORT")

//Represents an http server
type Server struct {

  ln net.Listener

  //handler to serve
  Handler *Handler

  //Bind address to open
  Addr string
}

//NewServer returns a new instance of Server
func NewServer() *Server {
  return &Server{
    Addr: DefaultAddr,
  }
}

//Open starts a http server
func (s *Server) Open() error {

  ln, err := net.Listen("tcp", s.Addr)
  if err != nil {
    return err
  }
  s.ln = ln

  //start a go routine that starts an http server
  go func() {http.Serve(s.ln, s.Handler)} ()

  return nil
}

//Close closes the socket
func (s *Server) Close() error {
  if s.ln != nil {
    s.ln.Close()
  }
  return nil
}

//Port returns the port the the HTTP server is currently running on -- only valid if a server is running. 
func (s *Server) Port() int {
  return s.ln.Addr().(*net.TCPAddr).Port
}

//Client represents a client to connect to the HTTP server
type Client struct {
  URL url.URL
  userService UserService
}

//NewClient returns a new instance of Client
func NewClient() *Client {
  c := &Client{}
  c.userService.URL = &c.URL
  return c
}

func (c *Client) UserService() tt.UserService {
  return &c.userService
}

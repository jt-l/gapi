package http_test

import (
  "github.com/james/TT/http"
)

type Handler struct {
  *http.Handler
  UserHandler *UserHandler
}

type Handler struct {
  *http.Handler

  DialHandler *DialHandler
}

func NewHandler() *Handler {
  h := &Handler {
    Handler: &http.Handler{},
    DialHandler: NewDialHandler(),
  }
  h.Handler.DialHandler = h.DialHandler.DialHandler
  return h
}

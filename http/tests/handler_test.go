package http_test

import (
  "github.com/james/TT/http"
)

type Handler struct {
  *http.Handler
  UserHandler *UserHandler
}

func NewHandler() *Handler {
  h := &Handler {
    Handler: &http.Handler{},
    UserHandler: NewUserHandler(),
  }
  h.Handler.UserHandler = h.UserHandler.UserHandler
  return h
}

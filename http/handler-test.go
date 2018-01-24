package http_test

import (
  "github.com/benbjohnson/wtf/http"
)

type Handler struct {
  *http.Handler

  UserHandler *UserHandler

}

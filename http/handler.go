package http


import (
  "encoding/json"
  "log"
  "net/http"
  "strings"

  "github.com/james/TT"
)

const ErrInvalidJSON = tt.Error("invalid json")

//Handler is a collection of all service handlers
type Handler struct {
  UserHandler *UserHandler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if strings.HasPrefix(r.URL.Path, "/api/user") {
    h.UserHandler.ServeHTTP(w,r)
  } else {
    http.NotFound(w, r)
  }
}

func Error(w http.ResponseWriter, err error, code int, logger *log.Logger) {
  //log the error
  logger.Printf("http error: %s (code=%d)", err, code)
  w.WriteHeader(code)
  json.NewEncoder(w).Encode(&errorResponse{Err: err.Error()})
}


type errorResponse struct {
  Err string `json:"err,omitempty"`
}

func NotFound(w http.ResponseWriter) {
  w.WriteHeader(http.StatusNotFound)
  w.Write([]byte(`{}` + "\n"))
}

// encodeJSON encodes v to w in JSON format. Error() is called if encoding fails.
func encodeJSON(w http.ResponseWriter, v interface{}, logger *log.Logger) {
	if err := json.NewEncoder(w).Encode(v); err != nil {
	Error(w, err, http.StatusInternalServerError, logger)
	}
}

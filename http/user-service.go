package http

import (
  "github.com/james/TT"
  "github.com/julienschmidt/httprouter"
  "log"
  "encoding/json"
  "os"
  "net/http"
  "net/url"
  "bytes"
)

//UserHandler represents an HTTP handler for users
type UserHandler struct {

  *httprouter.Router

  UserService tt.UserService

  Logger *log.Logger
}

//NewUserHandler creates an instance of UserHandler
func NewUserHandler() *UserHandler {
  h := &UserHandler{
    Router: httprouter.New(),
    Logger: log.New(os.Stderr, "", log.LstdFlags),
  }
  h.POST("/api/user", h.handlePostUser)
  return h
}

//handlePostUser handles requests to create a new user
func (h *UserHandler) handlePostUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  //Decode request
  var req postUserRequest
  if err := json.NewDecoder(r.Body).Decode(&req.User); err != nil {
    Error(w, ErrInvalidJSON, http.StatusBadRequest, h.Logger)
    return
  }
  u := req.User

  //create a new user
  err := h.UserService.CreateUser(u)
  if err != nil {
    Error(w, err, http.StatusBadRequest, h.Logger)
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(&postUserResponse{User: u})
}

type postUserRequest struct {
  User *tt.User `json:"User, omitempty"`
  Token string `json:"token, omitempty"`
}

type postUserResponse struct {
  User *tt.User `json:"User, omitempty"`
  Err string `json:"Err, omitempty"`
}

var _ tt.UserService = &UserService{}

type UserService struct {
  URL *url.URL
}

func (s *UserService) CreateUser(user *tt.User) error {
  if user == nil {
    return tt.ErrUserRequired
  }
  u := *s.URL
  u.Path = "/api/user/"

  token := user.Token
  name := user.Name

  reqBody, err := json.Marshal(postUserRequest{User: user, Token: token})
  if err != nil {
    return err
  }

  resp, err := http.Post(u.String(), "application/json", bytes.NewReader(reqBody))
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  var respBody postUserResponse
  if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
    return err
  } else if respBody.Err != "" {
    return tt.Error(respBody.Err)
  }
  *user = *respBody.User
  user.Name = name
  user.Token = token
  return nil
}

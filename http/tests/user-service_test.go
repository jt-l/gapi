package http_test

import (
  "github.com/james/TT/mock"
  "github.com/james/TT/http"
  "github.com/james/TT"

  "testing"
  "bytes"
  "log"
  "reflect"
)

func TestUserService_CreateUser(t *testing.T) {
  t.Run("OK", testUserService_CreateUser)
}

func testUserService_CreateUser(t *testing.T) {
  s, c := MustOpenServerAndClient()
  defer s.Close()

  //Mock service
  s.Handler.UserHandler.UserService.CreateUserFn = func(user *tt.User) error {
    if !reflect.DeepEqual(user, &tt.User{Name: "name", Token: "token"}) {
    }
    return nil
  }
  user := &tt.User{Name:"name", Token: "token"}
  err := c.UserService().CreateUser(user)
  if err != nil {
    t.Fatal(err)
  } else if !reflect.DeepEqual(user, &tt.User{Name: "name", Token: "token"}) {
    t.Fatalf("unexpected user: %#v", user)
  }
}

type UserHandler struct {
  *http.UserHandler
  UserService mock.UserService
  LogOutput bytes.Buffer
}

func NewUserHandler() *UserHandler {
  h := &UserHandler{UserHandler: http.NewUserHandler()}
  h.UserHandler.UserService = &h.UserService
  h.Logger = log.New(VerboseWriter(&h.LogOutput), "", log.LstdFlags)
  return h
}


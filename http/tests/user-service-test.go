package http_test

import (
//  "github.com/james/TT/"
  "testing"
)

func TestUserService_CreateUser(t *testing.T) {
  t.Run("OK", testUserService_CreateUser)
}

func testUserService_CreateUser(t *testing.T) {
  s, c := MustOpenServerAndClient()
  defer s.Close()

}


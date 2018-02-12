package mock

import (
	"github.com/james/TT"
)

type UserService struct {
	CreateUserFn      func(user *tt.User) error
	CreateUserInvoked bool
}

func (s *UserService) CreateUser(user *tt.User) error {
	s.CreateUserInvoked = true
	return s.CreateUserFn(user)
}

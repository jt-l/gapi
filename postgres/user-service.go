package postgres

import (
  "database/sql"
  "github.com/james/TT"
  _"github.com/lib/pq"
)

//UserService provides an implementation of the UserService interface using postgres
type UserService struct {

  DB *sql.DB

}

//Open opens a new connection to a postgres database
func Open(dataSourceName string) (*sql.DB, error) {
    db, err := sql.Open("postgres", dataSourceName)
    if err != nil {
        return nil, err
    }
    return db, nil
}

//CreateUser creates a new user
func (s *UserService) CreateUser(user *tt.User) error {
    stmt := "INSERT INTO users(name, address, phone, email, payperhour, hoursworked, timein, timeout, state, role, hashedpw, token, expire) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13);"
    _, err := s.DB.Exec(stmt, user.Name, user.Address, user.Phone, user.Email, user.PayPerHour, user.HoursWorked, user.TimeIn, user.TimeOut, user.State, user.Role, user.HashedPW, user.Token, user.Expire)
    if err != nil {
      return err
    }
    return nil
}


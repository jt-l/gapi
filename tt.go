package tt

type User struct {
  UserID int
  Name string
  Address string
  Phone string
  Email string
  PayPerHour int
  HoursWorked int
  TimeIn int
  TimeOut int
  State int
  Role int
  HashedPW string
  Token string
  Expire int64
}


type UserService interface {
  CreateUser(u *User) error
}

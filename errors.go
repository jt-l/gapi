package tt

//General errors

const (
  ErrUnauthorized = Error("unauthorized")
)

//User errors
const (
  ErrUserRequired = Error("User required")
)
//Error represents a tt error
type Error string

//Error returns a error message
func (e Error) Error() string {return string(e) }

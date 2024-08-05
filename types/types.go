package types 

// Interface because it is simpler to test interfaces in go
type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreatedUser(User) error
}


type User struct {
	ID int `json:"id"`
	FIrstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Password string `json:"_"`
	CreatedAt string `json:"createdAt"`
}

type ResgisterPayload struct {
	FIrstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Password string `json:"password"`

}
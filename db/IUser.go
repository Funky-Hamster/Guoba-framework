package db

type IUser interface {
	AddUser(u *User) (int64, error)

	UpdateUser(u *User) (int64, error)

	DeleteUserById(id int) (int64, error)

	GetUserByToken(token string) (*User, error)
}

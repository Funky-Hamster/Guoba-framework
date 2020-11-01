package usermodel

type IUser interface {
	// AddUser(u *User) (int64, error)

	// UpdateUser(u *User) (int64, error)

	// DeleteUserById(id int) (int64, error)

	GetUserByOpenid(jsCode string) (*User, error)
}

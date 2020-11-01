package usercontroller

import usermodel "guoba.io/guobastreet/models/user"

func login(jsCode string) (*usermodel.User, error) {
	var userOperation usermodel.IUser = usermodel.NewUserOperation()
	return userOperation.GetUserByOpenid(jsCode)
}

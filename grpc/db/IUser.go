package db

/**
 @func 用户操作接口
 @author:柠檬191104
 @date:2020/04/26
 */


type IUser interface {
	/**
	 *添加用户
	 */
	AddUser(u *User) (int64,error)
	/**
	 *更新用户
	 */
	UpdateUser(u *User) (int64,error)
	/**
	 根据ID删除用户
	 */
	DelUserById(id int) (int64,error)

	/**
	根据ID查询用户
	 */
	QueryUserById(id int) (*User,error)

	/**
	根据用户名查询用户
	@return 用户对象集
	 */
	QueryUserByName(username string)([]*User,error)
}

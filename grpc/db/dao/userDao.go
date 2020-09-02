package dao

import (
	"github.com/jmoiron/sqlx"
	"restfulServer/db"
	"restfulServer/db/conn"
)

/**
@func IUser接口的实现
@author:柠檬191104
@date:2020/04/26
*/
var dbConn *sqlx.DB

func init()  {
	dbConn  = conn.NewDb()
}
type UserDao struct {
}

func NewUserDao() *UserDao{
	return &UserDao{}
}



/**
 *添加用户
 */
func (user *UserDao) AddUser(u *db.User) (int64,error) {
	insert := "INSERT INTO user_tb(name,token) VALUES(?,?)"
	stmt,err := dbConn.Prepare(insert)
	if err != nil {
		return 0,err
	}
	defer conn.CloseStmt(stmt)
	result,err := stmt.Exec(u.Username,u.Password,u.Age,u.Sex)
	if err != nil {
		return 0,err
	}
	rowId,err := result.LastInsertId()
	if err != nil {
		return 0,err
	}

	return rowId,nil
}

/**
 *更新用户
 */
func (user *UserDao) UpdateUser(u *db.User) (int64,error) {
	update := "UPDATE user_tb SET name=?,token=? WHERE id=?"
	stmt,err := dbConn.Prepare(update)
	if err != nil {
		return  0,err
	}
	defer conn.CloseStmt(stmt)
	result, err := stmt.Exec(u.Username,u.Password,u.Age,u.Sex,u.Id)
	if err != nil {
		return 0,err
	}
	affecteRow,err := result.RowsAffected()
	if err != nil {
		return 0,err
	}
	return affecteRow,nil
}

/**
根据ID删除用户
*/
func (user *UserDao) DelUserById(id int) (int64,error) {

	delete := "DELETE FROM user_tb where id=?"
	result,err := dbConn.Exec(delete,id)
	if err != nil {
		return  0,err
	}
	affectedRow, err:= result.RowsAffected()
	if err !=nil {
		return 0,err
	}
	return affectedRow,nil
}

/**
根据token查询用户
*/
func (user *UserDao) QueryUserByToken(id int)(*db.User,error) {

	selectSql := "SELECT * FROM user_tb where token=?"
	var us *db.User = &db.User{}
	err := dbConn.Get(us,selectSql,id)
	if err != nil {
		return  nil,err
	}
	return  us,nil

}
package dao

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin/examples/grpc/db"
	"github.com/gin-gonic/gin/examples/grpc/db/conn"
	"github.com/jmoiron/sqlx"
)

var dbConn *sqlx.DB

func init() {
	dbConn = conn.NewDb()
}

type UserDao struct {
}

type code2SessionResponse struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Code       int    `json:"errcode"`
	Msg        string `json:"errmsg"`
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

/**
 *添加用户
 */
func (user *UserDao) AddUser(u *db.User) (int64, error) {
	insert := "INSERT INTO user_tb(name,token) VALUES(?,?)"
	stmt, err := dbConn.Prepare(insert)
	if err != nil {
		return 0, err
	}
	defer conn.CloseStmt(stmt)
	result, err := stmt.Exec(u.SessionKey, u.Openid)
	if err != nil {
		return 0, err
	}
	rowId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return rowId, nil
}

/**
 *更新用户
 */
func (user *UserDao) UpdateUser(u *db.User) (int64, error) {
	update := "UPDATE user_tb SET name=?,token=? WHERE id=?"
	stmt, err := dbConn.Prepare(update)
	if err != nil {
		return 0, err
	}
	defer conn.CloseStmt(stmt)
	result, err := stmt.Exec(u.SessionKey, u.Openid, u.Id)
	if err != nil {
		return 0, err
	}
	affecteRow, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affecteRow, nil
}

/**
根据token删除用户
*/
func (user *UserDao) DelUserById(id int) (int64, error) {

	deleteSql := "DELETE FROM user_tb where id=?"
	result, err := dbConn.Exec(deleteSql, id)
	if err != nil {
		return 0, err
	}
	affectedRow, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affectedRow, nil
}

/**
根据token查询用户
*/
func (u *UserDao) GetUserByToken(jsCode string) (*db.User, error) {

	// selectSql := "SELECT * FROM user_tb where token=?"
	var user *db.User = &db.User{}
	// err := dbConn.Get(user, selectSql, jsCode)
	// // Not found
	// if err != nil {
	resp, err := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=wx8f2e151800d8b54d&secret=2c1968bdc63021a0a063d73acc5d66a9&js_code=" + jsCode + "&grant_type=authorization_code")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var res code2SessionResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	user.Openid = res.Openid
	user.SessionKey = res.SessionKey

	if res.Code != 0 {
		return nil, nil
	}
	// u.AddUser(user)
	// }

	return user, nil

}

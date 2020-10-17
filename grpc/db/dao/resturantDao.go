package dao

import (
	"github.com/gin-gonic/gin/examples/grpc/db"
)

type ResturantDao struct {
}

func NewResturantDao() *ResturantDao {
	return &ResturantDao{}
}

func (r *ResturantDao) GetResturants() ([]*db.Resturant, error) {

	selectSql := "SELECT * FROM resturant_tb"
	var resturants []*db.Resturant
	err := dbConn.Select(&resturants, selectSql)
	// Not found
	if err != nil {
		return nil, err
	}

	return resturants, nil

}

func (resturant *ResturantDao) AddResturant(r *db.Resturant) (int64, error) {
	insert := "INSERT INTO resturant_tb(,) VALUES(?,?)"
	_, err := dbConn.Prepare(insert)
	if err != nil {
		return 0, err
	}
	// defer conn.CloseStmt(stmt)
	// result, err := stmt.Exec(u.SessionKey, u.Openid)
	// if err != nil {
	// 	return 0, err
	// }
	// rowId, err := result.LastInsertId()
	// if err != nil {
	// 	return 0, err
	// }
	return 0, nil
}

// 	return rowId, nil
// }

// /**
//  *更新用户
//  */
// func (resturant *ResturantDao) UpdateResturant(r *db.Resturant) (int64, error) {
// 	update := "UPDATE user_tb SET name=?,token=? WHERE id=?"
// 	stmt, err := dbConn.Prepare(update)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer conn.CloseStmt(stmt)
// 	result, err := stmt.Exec(u.SessionKey, u.Openid, u.Id)
// 	if err != nil {
// 		return 0, err
// 	}
// 	affecteRow, err := result.RowsAffected()
// 	if err != nil {
// 		return 0, err
// 	}
// 	return affecteRow, nil
// }

// /**
// 根据token删除用户
// */
// func (resturant *ResturantDao) DeleteResturantById(id int) (int64, error) {

// 	deleteSql := "DELETE FROM user_tb where id=?"
// 	result, err := dbConn.Exec(deleteSql, id)
// 	if err != nil {
// 		return 0, err
// 	}
// 	affectedRow, err := result.RowsAffected()
// 	if err != nil {
// 		return 0, err
// 	}
// 	return affectedRow, nil
// }

// /**
// 根据token查询用户
// */
// func (u *UserDao) GetResturantById(id int) (*db.Resturant, error) {

// 	// selectSql := "SELECT * FROM user_tb where token=?"
// 	var user *db.User = &db.User{}
// 	// err := dbConn.Get(user, selectSql, jsCode)
// 	// // Not found
// 	// if err != nil {
// 	resp, err := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=wx8f2e151800d8b54d&secret=2c1968bdc63021a0a063d73acc5d66a9&js_code=" + jsCode + "&grant_type=authorization_code")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	var res code2SessionResponse
// 	err = json.Unmarshal(body, &res)
// 	if err != nil {
// 		return nil, err
// 	}
// 	user.Openid = res.Openid
// 	user.SessionKey = res.SessionKey

// 	if res.Code != 0 {
// 		return nil, nil
// 	}
// 	// u.AddUser(user)
// 	// }

// 	return user, nil

// }

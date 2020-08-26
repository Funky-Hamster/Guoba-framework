package httpServer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"restfulServer/db"
	"restfulServer/db/dao"
	"strconv"
)

/**
@func 处理http请求
@author:柠檬191104
@date:2020/04/26
*/

type userParams struct{
	User db.User `json:"user"`
}

type myHttpServer struct {
}

func NewMyHttpServer() *myHttpServer {
	return &myHttpServer{}
}

var userDao db.IUser = dao.NewUserDao()

func (h *myHttpServer) StartServer(server string) {


	engine := gin.Default()
	//查询用户,可以根据用户名或id来查询
	engine.GET("/user/dao", func(c *gin.Context) {

		id := c.Query("id")
		if id != "" {
			idValue, err := strconv.Atoi(id)
			if err != nil {
				idValue = 0
			}
			var users []*db.User
			us, err := userDao.QueryUserById(idValue)
			if err != nil {
				users = append(users, nil)
			} else {
				users = append(users, us)
			}
			c.JSON(200, gin.H{
				"status": "ok",
				"users":  users,
				"error":  err,
			})
		} else {
			username := c.Query("username")
			us, err := userDao.QueryUserByName(username)
			c.JSON(200, gin.H{
				"status": "ok",
				"users":  us,
				"error":  err,
			})
		}
	})

	//处理添加用户请求
	engine.POST("/user/dao", func(c *gin.Context) {
		c.Header("content-Type","application/json")
		var user *db.User = &db.User{}
		err := c.BindJSON(user)
		fmt.Printf("get user data:%#v\n",user)
		if err != nil {
			c.JSON(200,gin.H{
				"status":"failed",
				"error":err,
			})
			return
		}
		rowId,err := userDao.AddUser(user)
		c.JSON(200,gin.H{
			"status":"ok",
			"error":err,
			"rowId":rowId,
		})
	})

	//处理更新用户请求
	engine.PUT("/user/dao", func(c *gin.Context) {
		c.Header("content-Type","application/json")
		var user *db.User = &db.User{}
		err := c.BindJSON(user)
		if err != nil {
			c.JSON(200,gin.H{
				"status":"failed",
				"error":err,
			})
			return
		}
		rowId,err := userDao.UpdateUser(user)
		c.JSON(200,gin.H{
			"status":"ok",
			"error":err,
			"affectedRow":rowId,
		})
	})

	engine.DELETE("/user/dao", func(c *gin.Context) {
		idStr := c.Query("id")
		if idStr == ""{
			idStr = "0"
		}
		id,err := strconv.Atoi(idStr)
		if err != nil {
			id = 0
		}
		afftecRow,err := userDao.DelUserById(id)
		c.JSON(200,gin.H{
			"status":"ok",
			"error":err,
			"affectedRow":afftecRow,
		})
	})

	engine.Run(server) //启动监听程序
}

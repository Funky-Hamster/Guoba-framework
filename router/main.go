package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	pb "github.com/gin-gonic/gin/examples/grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewSearchUserServiceClient(conn)

	// Set up a http server.
	r := gin.Default()
	r.GET("/guoba/user/:token", func(c *gin.Context) {
		token := c.Param("token")

		// Contact the server and print out its response.
		req := &pb.SearchRequest{Openid: token}
		res, err := client.Search(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"data": nil,
				"msg":  err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": res.Code,
			"data": res.Data,
			"msg":  res.Msg,
		})
	})

	r.GET("/guoba/resturant", func(c *gin.Context) {
		client := pb.NewListRestaurantsServiceClient(conn)
		// Contact the server and print out its response.
		req := &pb.ListRestaurantsRequest{Limit: 1, Page: 1}
		res, err := client.List(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"data": nil,
				"msg":  err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": res.Code,
			"data": res.Data,
			"msg":  res.Msg,
		})
	})

	// r.POST("/guoba/user", func(c *gin.Context) {
	// 	//var user *db.User = &db.User{}
	// 	//var user db.User
	// 	user := &db.User{SessionKey: c.PostForm("session_key"), Openid: c.PostForm("openid")}
	// 	fmt.Printf("get user data:%#v\n", user)
	// 	req := &pb.AddUserRequest{Openid: user.Openid, SessionKey: user.SessionKey}
	// 	res, err := client.AddUser(c, req)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{
	// 			"code": 500,
	// 			"data": nil,
	// 			"msg":  err.Error(),
	// 		})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code": res.Code,
	// 		"data": res.Data,
	// 		"msg":  res.Msg,
	// 	})
	// })

	// Run http server
	if err := autotls.Run(r, "guoba.online", "localhost", "www.guoba.online"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}

}

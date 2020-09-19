package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/examples/grpc/db"
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
		req := &pb.SearchRequest{Token: token}
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

	r.POST("/guoba/user", func(c *gin.Context) {
		//var user *db.User = &db.User{}
		//var user db.User
		user := &db.User{Name: c.PostForm("name"), Token: c.PostForm("token")}
		fmt.Printf("get user data:%#v\n", user)
		req := &pb.AddUserRequest{Token: user.Token, Name: user.Name}
		res, err := client.AddUser(c, req)
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

	// Run http server
	if err := r.Run(":8052"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

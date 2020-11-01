package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set up a HTTPS server.
	r := gin.Default()
	r.GET("/guoba/user/:openid", func(c *gin.Context) {
		openid := c.Param("openid")

		// Contact the server and print out its response.
		req := &pb.SearchRequest{Openid: openid}
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

	// Run http server
	if err := autotls.Run(r, "guoba.online", "localhost", "www.guoba.online"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}

}

package main

import (
	"restfulServer/httpServer"
)

func main(){

	myHttpServer := httpServer.NewMyHttpServer()

	myHttpServer.StartServer(":8080")
}

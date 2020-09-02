package main

import (
	"log"
	"net"

	pb "github.com/gin-gonic/gin/examples/grpc/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server is used to implement helloworld.GreeterServer.
type userServer struct{}

func (s *userServer) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchResponse, error) {
	// Add DB connection here
	// QueryUserById(1)
	return &pb.SearchResponse{Code: 200, Data: nil, Msg: ""}, nil
}
func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSearchUserServiceServer(s, &userServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


// func (user *UserDao) QueryUserById(id int)(*db.User,error) {

// 	selectSql := "SELECT * FROM user_tb where id=?"
// 		var us *db.User = &db.User{}
// 		err := dbConn.Get(us,selectSql,id)
// 		if err != nil {
// 			return  nil,err
// 		}
// 		return  us,nil

// }

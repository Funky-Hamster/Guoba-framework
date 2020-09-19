package main

import (
	"log"
	"net"

	db "github.com/gin-gonic/gin/examples/grpc/db"
	dao "github.com/gin-gonic/gin/examples/grpc/db/dao"
	pb "github.com/gin-gonic/gin/examples/grpc/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server is used to implement helloworld.GreeterServer.
type userServer struct{}

var userDao db.IUser = dao.NewUserDao()

func (s *userServer) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchResponse, error) {
	var user *db.User
	user, err := userDao.GetUserByToken(in.Token)
	if err != nil {
		return &pb.SearchResponse{Code: 500, Data: nil, Msg: err.Error()}, err
	}
	if user == nil {
		return &pb.SearchResponse{Code: 404, Data: nil, Msg: "Not found"}, nil
	}
	return &pb.SearchResponse{Code: 200, Data: &pb.User{Id: user.Id, Name: user.Name, Token: user.Token}, Msg: ""}, nil
}

func (s *userServer) AddUser(ctx context.Context, in *pb.AddUserRequest) (*pb.SearchResponse, error) {
	user := &db.User{Name: in.Name, Token: in.Token}
	rowID, err := userDao.AddUser(user)
	if err != nil {
		return &pb.SearchResponse{Code: 500, Data: nil, Msg: err.Error()}, err
	}
	if user == nil {
		return &pb.SearchResponse{Code: 404, Data: nil, Msg: "Not found"}, nil
	}
	return &pb.SearchResponse{Code: 200, Data: &pb.User{Id: int32(rowID), Name: user.Name, Token: user.Token}, Msg: ""}, nil
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

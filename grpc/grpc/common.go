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

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSearchUserServiceServer(s, &userServer{})
	pb.RegisterListRestaurantsServiceServer(s, &resturantServer{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type userServer struct{}

var userDao db.IUser = dao.NewUserDao()

func (s *userServer) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchResponse, error) {
	var user *db.User
	user, err := userDao.GetUserByToken(in.Openid)
	if err != nil {
		return &pb.SearchResponse{Code: 500, Data: nil, Msg: err.Error()}, err
	}
	if user == nil {
		return &pb.SearchResponse{Code: 404, Data: nil, Msg: "Not found"}, nil
	}
	return &pb.SearchResponse{Code: 200, Data: &pb.User{Id: user.Id, SessionKey: user.SessionKey, Openid: user.Openid}, Msg: ""}, nil
}

func (s *userServer) AddUser(ctx context.Context, in *pb.AddUserRequest) (*pb.SearchResponse, error) {
	user := &db.User{SessionKey: in.SessionKey, Openid: in.Openid}
	rowID, err := userDao.AddUser(user)
	if err != nil {
		return &pb.SearchResponse{Code: 500, Data: nil, Msg: err.Error()}, err
	}
	if user == nil {
		return &pb.SearchResponse{Code: 404, Data: nil, Msg: "Not found"}, nil
	}
	return &pb.SearchResponse{Code: 200, Data: &pb.User{Id: int32(rowID), SessionKey: user.SessionKey, Openid: user.Openid}, Msg: ""}, nil
}

type resturantServer struct{}

var resturantDao db.IResturant = dao.NewResturantDao()

func (s *resturantServer) List(ctx context.Context, in *pb.ListRestaurantsRequest) (*pb.ListRestaurantsResponse, error) {
	var resturants []*db.Resturant
	resturants, err := resturantDao.GetResturants()
	var resResturants []*pb.Restaurant = make([]*pb.Restaurant, 0)
	for i := range resturants {
		resResturants = append(resResturants, &pb.Restaurant{Id: resturants[i].Id, Name: resturants[i].Name, Description: resturants[i].Description, Discounts: resturants[i].Discounts, ImgUrl: resturants[i].ImgUrl})
	}
	if err != nil {
		return &pb.ListRestaurantsResponse{Code: 500, Data: nil, Msg: err.Error()}, err
	}
	if resturants == nil {
		return &pb.ListRestaurantsResponse{Code: 404, Data: nil, Msg: "Not found"}, nil
	}
	return &pb.ListRestaurantsResponse{Code: 200, Data: resResturants, Msg: ""}, nil
}

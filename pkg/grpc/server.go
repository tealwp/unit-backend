package grpc

import (
	"context"
	"log"
	"net"

	"github.com/tealwp/unit-backend/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedInventoryServer
}

func (s *server) GetItemList(ctx context.Context, in *pb.GetItemListRequest) (*pb.GetItemListResponse, error) {
	return &pb.GetItemListResponse{
		Items: getSampleItems(),
	}, nil
}

func Serve() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterInventoryServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getSampleItems() []*pb.Item {
	return []*pb.Item{
		{
			Name: "testing",
		},
		{
			Name: "testing2",
		},
		{
			Name: "testing3",
		},
	}
}

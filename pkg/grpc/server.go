package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/tealwp/unit-backend/pkg/cfg"
	"github.com/tealwp/unit-backend/pkg/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedUnitServer
}

func (s *server) GetEventList(ctx context.Context, in *pb.GetEventListRequest) (*pb.GetEventListResponse, error) {
	return &pb.GetEventListResponse{
		Events: getSampleEvents(),
	}, nil
}

func (s *server) GetEvent(ctx context.Context, in *pb.GetEventRequest) (*pb.GetEventResponse, error) {
	return &pb.GetEventResponse{
		Event: getSampleEvent(in.Id),
	}, nil
}

func (s *server) CreateEvent(ctx context.Context, in *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
	return &pb.CreateEventResponse{}, nil
}

func Serve(cfg *cfg.Config) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC_PORT))
	if err != nil {
		return fmt.Errorf("creating listener: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUnitServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}

func getSampleEvents() []*pb.Event {
	return []*pb.Event{
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

func getSampleEvent(id int64) *pb.Event {
	return &pb.Event{
		Name: fmt.Sprintf("testing_event %d", id),
	}
}

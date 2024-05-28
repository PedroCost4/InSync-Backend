package server

import (
	"net"

	"github.com/PedroCost4/movies-backend/config"
	"github.com/PedroCost4/movies-backend/handler"
	"github.com/PedroCost4/movies-backend/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func InitServer() error {
	logger := config.GetLogger("main")

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		logger.Errorf("Unable to start the server: %v\n", err)
		return err
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterUsersServer(s, &handler.UserServer{})

	if err := s.Serve(listener); err != nil {
		logger.Errorf("Unable to start the server: %v\n", err)
		return err
	}

	return nil
}

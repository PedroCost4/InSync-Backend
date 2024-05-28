package handler

import (
	"context"

	"github.com/PedroCost4/movies-backend/pb"
	"github.com/PedroCost4/movies-backend/schema"
)

type UserServer struct {
	pb.UnimplementedUsersServer
}

func (s *UserServer) GetUser(ctx context.Context, userReq *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user := schema.User{}
	err := db.Get(&user, "SELECT * FROM users WHERE id = $1", userReq.Id)
	if err != nil {
		logger.Errorf("Unable to get user: %v\n", err)
		return nil, err
	}

	returnedUser := &pb.User{
		Id:    user.ID.String(),
		Name:  user.Username,
		Email: user.Email,
	}

	return &pb.GetUserResponse{
		User: returnedUser,
	}, nil

}

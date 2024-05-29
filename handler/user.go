package handler

import (
	"context"

	"github.com/PedroCost4/movies-backend/pb"
	"github.com/PedroCost4/movies-backend/schema"
	"github.com/PedroCost4/movies-backend/utils"
	"github.com/google/uuid"
)

type UserServer struct {
	pb.UnimplementedUsersServer
}

func (s *UserServer) GetUser(ctx context.Context, userReq *pb.GetUserRequest) (*pb.UserResponse, error) {
	user := schema.User{}
	err := db.Get(&user, "SELECT * FROM users WHERE id = $1", userReq.Id)
	if err != nil {
		logger.Errorf("Unable to get user: %v\n", err)
		return nil, err
	}

	returnedUser := &pb.User{
		Id:       user.ID.String(),
		UserName: user.Username,
		Email:    user.Email,
	}

	return &pb.UserResponse{
		User: returnedUser,
	}, nil

}

func (s *UserServer) Register(ctx context.Context, userReq *pb.CreateUserRequest) (*pb.UserResponse, error) {
	hashedPassword, err := utils.HashPassword(userReq.Password)

	if err != nil {
		logger.Errorf("Unable to hash password: %v\n", err)
		return nil, err
	}

	newUUID := uuid.New()

	id := uuid.UUID{}
	err = db.QueryRowx(`INSERT INTO
		users (id, username, email, password)
	    VALUES ($1, $2, $3, $4)
		RETURNING id`, newUUID.String(), userReq.UserName, userReq.Email, hashedPassword).Scan(&id)

	if err != nil {
		logger.Errorf("Unable to create user: %v\n", err)
		return nil, err
	}

	returnedUser := &pb.User{
		Id:       id.String(),
		UserName: userReq.UserName,
		Email:    userReq.Email,
	}

	return &pb.UserResponse{
		User: returnedUser,
	}, nil
}

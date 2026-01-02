package handlers

import (
	"context"

	"github.com/adarshjeetAmplio/grpc-server/internal/services"
	proto "github.com/adarshjeetAmplio/grpc-server/proto"
)

type GrpcServer struct {
	proto.UnimplementedUserServiceServer
	userService services.IUserService
}

func NewUserGRPCServer(userService services.IUserService) *GrpcServer{
	return &GrpcServer{
		userService: userService,
	}
}

func (gs *GrpcServer) Signup(ctx context.Context, in *proto.SignupRequest) (*proto.SignupResponse, error){
	resp, err:= gs.userService.Signup(in)
	return resp, err
}
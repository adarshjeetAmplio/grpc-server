package services

import (
	"github.com/adarshjeetAmplio/grpc-server/internal/dal"
	"github.com/adarshjeetAmplio/grpc-server/internal/data/models"
	"github.com/adarshjeetAmplio/grpc-server/internal/utils"
	proto "github.com/adarshjeetAmplio/grpc-server/proto"
	"gorm.io/gorm"
	
)

type IUserService interface{
	Signup(*proto.SignupRequest) (*proto.SignupResponse,error)
}

type userService struct {
	userDal dal.IUserDal
}

func (u *userService) Signup(in *proto.SignupRequest) (*proto.SignupResponse, error){
	db:= utils.GetDB()
	user := &models.User{
		Email:    in.Email,
		Name:     in.Name,
		Password: in.Password,
	}
	err:= InsertUser(db, user)
	if err != nil {
		return &proto.SignupResponse{Message: "failed to create account", Token: "NA"}, err
	}
	return &proto.SignupResponse{Message: "Account Created Successfully", Token:"q345uiosdfghjwertyu"}, nil
}

func InsertUser(db *gorm.DB, record *models.User) error {
	if err := db.Create(record).Error; err != nil {
		return err
	}
	return nil
}

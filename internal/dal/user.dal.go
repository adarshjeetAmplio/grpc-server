package dal

import (
	"errors"
	"github.com/adarshjeetAmplio/grpc-server/internal/data/models"
	"gorm.io/gorm"
)


type IUserDal interface {
	// IBaseDal[models.User]
	InsertUser(db *gorm.DB, record *models.User) error
}

type UserDal struct{
	*DBContext
}


func NewUserDal(conn *gorm.DB) IUserDal {
	return &UserDal{
		DBContext: NewDBContext(conn),
	}
}

func (ud *UserDal) InsertUser(db *gorm.DB, record *models.User) error {
	result := ud.DBContext.Connection.Create(record)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("could not insert row in investor table")
	}
	return nil
}

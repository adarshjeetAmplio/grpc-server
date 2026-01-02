package dal

import "gorm.io/gorm"

type DBContext struct {
	Connection *gorm.DB
}

type IBaseDal[T any] interface {
	Insert(record *T) (string, error)
	Update(record *T) error
	GetById(id string) (*T, error)
	Get() ([]*T, error)
	Upsert(record *T) error
	Delete(id string) error
	SoftDelete(id string) error
}

func NewDBContext(conn *gorm.DB) *DBContext {
	return &DBContext{
		Connection: conn,
	}
}
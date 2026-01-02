package data

import (
	"os"

	"github.com/adarshjeetAmplio/grpc-server/internal/data/models"
	"github.com/adarshjeetAmplio/grpc-server/internal/utils"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	
	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USERNAME")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")
	db_schema := os.Getenv("DB_SCHEMA")
	db_password := os.Getenv("DB_PASSWORD")

	utils.InitializeDatabaseLayer(db_host, db_port, db_name, db_user, db_password, db_schema, getAutoMigrateModels())

	return nil;
}

func getAutoMigrateModels() []interface{}{
	return []any{
		models.User{},
	}
}
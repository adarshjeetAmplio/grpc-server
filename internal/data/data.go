package data

import (
	"fmt"
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

	if db_host == "" || db_user == "" || db_name == "" {
		err:= fmt.Errorf("missing required database environment variables")
		fmt.Println(err)
	}

	utils.InitializeDatabaseLayer(db_host, db_port, db_name, db_user, db_password, db_schema, getAutoMigrateModels())

	return nil;
}

func getAutoMigrateModels() []interface{}{
	return []any{
		models.User{},
	}
}
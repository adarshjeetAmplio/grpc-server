package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dbInstance *gorm.DB

// returns the connected database instance
func GetDB() *gorm.DB {
	return dbInstance
}

// initialize the database layer for the provided postgres database credentials and run the auto migrations passed in params
func InitializeDatabaseLayer(dbHost, dbPort, dbName, dbUser, dbPassword string, schema string, initSchema []interface{}) error {
	db, err := setupPostgres(dbHost, dbPort, dbName, dbUser, dbPassword, schema)
	if err != nil {
		return err
	}

	result := db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s;", schema))

	if result.Error != nil {
		return err
	}

	result = db.Exec(fmt.Sprintf("SET search_path = %s;", schema))

	if result.Error != nil {
		return err
	}

	// err = autoMigrate(db, migrations, initSchema)

	// if err != nil {
	// 	return err
	// }
	database, err := db.DB()
	if err != nil {
		return err
	}
	database.SetMaxIdleConns(1)
	database.SetConnMaxIdleTime(time.Hour * 12)
	database.SetMaxOpenConns(10)

	dbInstance = db
	return nil
}
func setupPostgres(dbHost, dbPort, dbName, dbUser, dbPassword string, dbSchema string ) (*gorm.DB, error) {
	connnectionString := fmt.Sprintf("host=%s post=%s user=%s dbname=%s password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)

	return gorm.Open(postgres.Open(connnectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: fmt.Sprintf("%s.", dbSchema),
			SingularTable: false,
		},
	})
}

func DbTable(db *gorm.DB, schema string, tableName string) *gorm.DB {
	return db.Scopes(_dbSchemaResolver(schema, tableName))
}

func _dbSchemaResolver(schema string, tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		tableName := schema + "." + tableName
		return db.Table(tableName)
	}
}

func autoMigrate(db *gorm.DB, migrations []*gormigrate.Migration, initSchema []interface{}) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrations)

	m.InitSchema(func(tx *gorm.DB) error {
		err := tx.AutoMigrate(initSchema...)
		if err != nil {
			return err
		}
		return nil
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
		return err
	}

	log.Printf("Migration did run successfully")
	return nil
}
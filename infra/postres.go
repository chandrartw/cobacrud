package infra

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/chandrartw/cobacrud/model"
)

func LoadPostgresSQLDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 user=postgres dbname=nprm-db password=admin sslmode=disable")
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&model.Contract_Accounts{})
	return db
}
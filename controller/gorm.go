package controller

import (
	"github.com/jinzhu/gorm"
	"github.com/chandrartw/cobacrud/infra"
)

type InDB struct {
	DB *gorm.DB
}

func ConnectDB() *InDB{
	db := infra.LoadPostgresSQLDB()
	return inDB
}
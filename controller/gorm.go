package controller

import (
	"github/jinzhu/gorm"
)

type InDB struct {
	DB *gorm.DB
}

func ConnectDB *InDB{
	db := infra.LoadPostgresSQLDB()
	return inDB
}
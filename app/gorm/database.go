package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CreateConnection() (*gorm.DB, error) {
	dsn := GetEnv()
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		return nil, err
	} else {
		return db, err
	}

}

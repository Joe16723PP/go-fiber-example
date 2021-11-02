package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CreateConnection() (*gorm.DB, error) {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname=testDB?charset=utf8mb4&parseTime=True&loc=Local"
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		return nil, err
	} else {
		return db, err
	}

}

package gorm

import "fmt"

type DatabaseEnv struct {
	User     string
	Password string
	DbName   string
	Port     string
}

var localEnv = DatabaseEnv{
	User:     "root",
	Password: "200119",
	DbName:   "testDb",
	Port:     "3306",
}

func GetEnv() string {
	return fmt.Sprintf("%v:%v@tcp(127.0.0.1:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		localEnv.User, localEnv.Password, localEnv.Port, localEnv.DbName)
}

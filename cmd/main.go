package main

import (
	"fmt"
	database "go-fiber-example/app/gorm" // gorm sql driver
	"go-fiber-example/app/server"
	vr "go-fiber-example/utils/validator" // validator
	"log"
)

func main() {
	var err error
	// init database
	database.DB, err = database.CreateConnection()
	// init go validator
	vr.Validator = vr.InitValidator()
	if err != nil {
		panic(fmt.Sprintf("create connection err: %v", err))
	}
	// init go fiber app
	server.App = server.InitApp()

	// listen on port 3000
	log.Fatalf(server.App.Listen(":3000").Error())
}

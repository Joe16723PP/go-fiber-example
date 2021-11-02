package main

import (
	"fmt"
	database "go-fiber-example/app/gorm" // gorm sql driver
	"go-fiber-example/app/server"
	"go-fiber-example/routes"
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
	app := server.InitApp()
	// grouping route
	apiV1 := app.Group("/v1") // versioning
	userRouter := apiV1.Group("/user")
	// register router
	routes.RegisterHealthCheckRouter(&apiV1)
	routes.RegisterUserRouter(&userRouter)

	// listen on port 3000
	log.Fatalf(app.Listen(":3000").Error())
}

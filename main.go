package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jonasssneto/go-crud/src/controller/routes"
	controller "github.com/jonasssneto/go-crud/src/controller/user"
	"github.com/jonasssneto/go-crud/src/model/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// init dependencies
	userService := service.NewUserDomainService()
	userController := controller.NewUserConctrollerInterface(userService)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}

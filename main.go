package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jonasssneto/go-crud/src/configuration/database/mongodb"
	"github.com/jonasssneto/go-crud/src/controller/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())

	if err != nil {
		log.Fatal("error to connect database", err.Error())
		return
	}

	router := gin.Default()

	userController := initDependencies(database)

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}

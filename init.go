package main

import (
	controller "github.com/jonasssneto/go-crud/src/controller/user"
	"github.com/jonasssneto/go-crud/src/model/repository"
	"github.com/jonasssneto/go-crud/src/model/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	userService := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(userService)
}

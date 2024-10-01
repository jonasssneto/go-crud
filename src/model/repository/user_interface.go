package repository

import (
	restErr "github.com/jonasssneto/go-crud/src/configuration/err"
	"github.com/jonasssneto/go-crud/src/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type userRepository struct {
	database *mongo.Database
}

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		database,
	}
}

type UserRepositoryInterface interface {
	Create(userDomain model.UserDomainInterface) *restErr.Err
}

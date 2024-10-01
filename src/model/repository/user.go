package repository

import (
	"context"
	"os"

	restErr "github.com/jonasssneto/go-crud/src/configuration/err"
	"github.com/jonasssneto/go-crud/src/configuration/logger"
	"github.com/jonasssneto/go-crud/src/model"
	"github.com/jonasssneto/go-crud/src/model/repository/entity/converter"
	"go.uber.org/zap"
)

type UserRepository interface {
	Create(userDomain model.UserDomainInterface) *restErr.Err
}

var (
	MONGO_COLLECTION = "MONGO_COLLECTION"
)

func (ur *userRepository) Create(userDomain model.UserDomainInterface) *restErr.Err {
	logger.Info("Init create repository")

	collection_name := os.Getenv(MONGO_COLLECTION)

	collection := ur.database.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	_, err := collection.InsertOne(context.Background(), value)

	if err != nil {
		logger.Error("Error trying to create user",
			err,
			zap.String("journey", "createUser"))
		return restErr.NewInternalServerError(err.Error())
	}

	logger.Info(
		"CreateUser repository executed successfully",
		zap.String("userId", value.ID.Hex()),
		zap.String("journey", "createUser"))

	return nil
}

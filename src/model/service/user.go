package service

import (
	restErr "github.com/jonasssneto/go-crud/src/configuration/err"
	"github.com/jonasssneto/go-crud/src/configuration/logger"
	"github.com/jonasssneto/go-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) Create(userDomain model.UserDomainInterface) *restErr.Err {
	logger.Info("Init CreateUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	err := ud.userRepository.Create(userDomain)

	if err != nil {
		return nil
	}

	return nil
}

func (ud *userDomainService) Update(id string, userDomain model.UserDomainInterface) *restErr.Err {
	logger.Info("Init UpdateUser model", zap.String("journey", "updateUser"))

	return nil
}

func (ud *userDomainService) FindById(id string) (*model.UserDomainInterface, *restErr.Err) {
	logger.Info("Init FindById model", zap.String("journey", "findUser"))

	return nil, nil
}

func (ud *userDomainService) Delete(id string) *restErr.Err {
	logger.Info("Init DeleteUser model", zap.String("journey", "deleteUser"))

	return nil
}

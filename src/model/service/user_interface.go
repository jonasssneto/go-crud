package service

import (
	restErr "github.com/jonasssneto/go-crud/src/configuration/err"
	"github.com/jonasssneto/go-crud/src/model"
	"github.com/jonasssneto/go-crud/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	Create(model.UserDomainInterface) *restErr.Err
	Update(string, model.UserDomainInterface) *restErr.Err
	FindById(string) (*model.UserDomainInterface, *restErr.Err)
	Delete(string) *restErr.Err
}

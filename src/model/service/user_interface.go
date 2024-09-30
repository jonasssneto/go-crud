package service

import (
	restErr "github.com/jonasssneto/go-crud/src/configuration/err"
	"github.com/jonasssneto/go-crud/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct{}

type UserDomainService interface {
	Create(model.UserDomainInterface) *restErr.Err
	Update(string, model.UserDomainInterface) *restErr.Err
	FindById(string) (*model.UserDomainInterface, *restErr.Err)
	Delete(string) *restErr.Err
}

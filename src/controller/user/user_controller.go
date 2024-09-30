package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jonasssneto/go-crud/src/model/service"
)

func NewUserConctrollerInterface(serviceInterface service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	FindById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}

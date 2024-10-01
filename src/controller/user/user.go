package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonasssneto/go-crud/src/configuration/logger"
	"github.com/jonasssneto/go-crud/src/configuration/validation"
	"github.com/jonasssneto/go-crud/src/controller/model/request"
	"github.com/jonasssneto/go-crud/src/model"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) FindById(c *gin.Context) {
}

func (uc *userControllerInterface) Create(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)

	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		err := validation.ValidateUserError(err)

		c.JSON(err.Code, err)
		return
	}

	domain := model.NewUserDomain(
		UserRequest.Email,
		UserRequest.Password,
		UserRequest.Name,
		UserRequest.Age,
	)

	err := uc.service.Create(domain)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"CreateUser controller executed successfully",
		zap.String("journey", "createUser"))

	c.JSON(http.StatusCreated, "")
}

func (uc *userControllerInterface) Update(c *gin.Context) {
}

func (uc *userControllerInterface) Delete(c *gin.Context) {
}

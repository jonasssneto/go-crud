package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/jonasssneto/go-crud/src/controller/user"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface,
) {
	r.POST("/user", userController.Create)
	r.GET("/user/:userId", userController.FindById)
	r.PUT("/user/:userId", userController.Update)
	r.DELETE("/user/:userId", userController.Delete)
}

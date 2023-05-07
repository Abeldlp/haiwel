package route

import (
	"github.com/Abeldlp/haiwel/user-service/controller"
	"github.com/gin-gonic/gin"
)

func InitializeUserRoutes(group *gin.RouterGroup) {
	users := group.Group("/users")

	users.GET("", controller.GetAllUsers)
}

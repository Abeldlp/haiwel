package route

import (
	"github.com/Abeldlp/haiwel/user-service/controller"
	"github.com/gin-gonic/gin"
)

func InitializeUserRankingRoutes(group *gin.RouterGroup) {
	usersRanking := group.Group("/user-rankings")

	usersRanking.GET("", controller.GetAllUserRankings)
}

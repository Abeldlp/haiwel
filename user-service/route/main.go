package route

import "github.com/Abeldlp/haiwel/user-service/config"

func InitializeRoutes() {

	// V1 ROUTES
	InitializeUserRoutes(config.V1)
	InitializeUserRankingRoutes(config.V1)
}

package config

import (
	"github.com/gin-gonic/gin"
)

var V1 *gin.RouterGroup

func InitializeHttpServer() *gin.Engine {
	router := gin.Default()

	V1 = router.Group("/api/v1")

	return router
}

package controller

import (
	"github.com/Abeldlp/haiwel/user-service/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUserRankings(c *gin.Context) {
	rankings, err := service.GetAllRankings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
	}

	c.JSON(http.StatusOK, rankings)
}

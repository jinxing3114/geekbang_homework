package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "获取用户信息",
	})
}
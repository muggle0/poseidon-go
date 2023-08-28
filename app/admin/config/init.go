package config

import (
	"github.com/gin-gonic/gin"
)

func InitMiddleware(c *gin.Engine) {
	// 数据库链接
	c.Use(WithContextDb)
}

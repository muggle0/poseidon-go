package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hahahahahhaahahahah",
		})
	})
	r.POST("/test1", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		query := context.Query("name")
		context.Writer.Write([]byte("message->" + query))
	})
	r.Run(":9090") // listen and serve on 0.0.0.0:8080
}

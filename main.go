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
		result, exit := context.GetPostForm("post")
		if exit {
			fmt.Println(result)
		}
		context.Writer.Write([]byte("message->" + query))
	})
	r.Handle("DELETE", "/test2/:id", func(context *gin.Context) {
		param := context.Param("id")
		context.Writer.Write([]byte(param))
	})
	r.Run(":9090") // listen and serve on 0.0.0.0:8080
}

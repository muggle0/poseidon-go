package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	/*import 下划线（如：import _ hello/imp）的作用：当导入一个包时，该包下的文件里所有init()函数都会被执行，
	然而，有些时候我们并不需要把整个包都导入进来，仅仅是是希望它执行init()函数而已。
	这个时候就可以使用 import _ 引用该包。即使用【import _ 包路径】只是引用该包，
	仅仅是为了调用init()函数，所以无法通过包名来调用包中的其他函数。*/
	"poseidon-go/Model"
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
	r.GET("/test3", func(context *gin.Context) {
		/**
		但是在go中，你没有明确提到一个类型是否实现了一个接口。 如果一个类型实现了在接口中定义的签名方法，则称该类型实现该接口。 接口就是定义了对象的行为。
		*/
		var user Model.User
		/**
		http://localhost:9090/test3?username=nih&phone=sss
		*/
		err := context.ShouldBindQuery(&user)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(user.Phone)
		context.Writer.Write([]byte(user.Phone))
	})
	r.POST("/test4", func(context *gin.Context) {
		var user Model.User
		context.BindJSON(&user)
		fmt.Println(user.Phone)
	})
	r.Run(":9090") // listen and serve on 0.0.0.0:8080
}

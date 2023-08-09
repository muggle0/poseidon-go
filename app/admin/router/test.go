package router

import (
	"github.com/gin-gonic/gin"
	"poseidon-go/app/admin/apis"
)

func init() {
	RouterList = append(RouterList, registerTestRouter)
}

func registerTestRouter(v1 *gin.RouterGroup) {
	api := apis.Test{}
	r := v1.Group("/test")
	{
		r.GET("", api.GetPage)

	}
}

package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
)

type Test struct {
	api.Api
}

// 函数名称前面的括号是Go定义这些函数将在其上运行的对象的方式
func (e Test) GetPage(c *gin.Context) {
	e.MakeContext(c)
	e.OK(">>>>", "xxxxxxxxxxxxx")
}

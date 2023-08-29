package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"poseidon-go/app/admin/model/dto"
	"poseidon-go/app/admin/service"
)

type Test struct {
	api.Api
}

// 函数名称前面的括号是Go定义这些函数将在其上运行的对象的方式
func (e Test) GetPage(c *gin.Context) {
	req := dto.TestReq{}
	// 结构体初始化
	s := service.Test{}
	// IOC 手动挡
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	err = s.Message(&req).Error
	e.OK(">>>>", "xxxxxxxxxxxxx")
}

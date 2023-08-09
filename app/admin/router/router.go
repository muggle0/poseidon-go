package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"os"
)

var RouterList = make([]func(*gin.RouterGroup), 0)

// InitRouter 路由初始化，不要怀疑，这里用到了
func InitRouter() {
	var r *gin.Engine
	h := sdk.Runtime.GetEngine()
	if h == nil {
		h = gin.New()
		sdk.Runtime.SetEngine(h)
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
		os.Exit(-1)
	}
	// 可根据业务需求来设置接口版本
	v1 := r.Group("/api/v1")
	for _, f := range RouterList {
		f(v1)
	}
	r.Run(":8080")
}

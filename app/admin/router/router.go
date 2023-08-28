package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"poseidon-go/app/admin/config"
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
	db, err := gorm.Open(mysql.Open("root:root@tcp/vcp_saas?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Error(err)
	}
	sdk.Runtime.SetDb("*", db)
	// 这个必须在路由组之前设置，否则无效
	config.InitMiddleware(r)
	// 可根据业务需求来设置接口版本
	v1 := r.Group("/api/v1")
	for _, f := range RouterList {
		f(v1)
	}
	r.Run(":8080")
}

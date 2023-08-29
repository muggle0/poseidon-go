package service

import (
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk"
	"gorm.io/gorm"
	"poseidon-go/app/admin/model/dto"

	"github.com/go-admin-team/go-admin-core/sdk/service"
)

type Test struct {
	service.Service
}

// GetPage 获取SysApi列表
func (e *Test) Message(req *dto.TestReq) *Test {
	db := sdk.Runtime.GetDbByKey("*")
	err := db.Model(&str).Scopes().
		First(model, req.GetType()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSysApi error:%s", err)
		_ = e.AddError(err)
		return e
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		_ = e.AddError(err)
		return e
	}
	return e
}

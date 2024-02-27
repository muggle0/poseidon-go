package service

import (
	"poseidon-go/app/admin/model/dto"

	"github.com/go-admin-team/go-admin-core/sdk/service"
)

// 此处类似autowrite 注入
type Test struct {
	service.Service
}

func (e *Test) Message(req *dto.TestReq) *Test {
	deptList := make([]int, 0)
	if err := e.Orm.Table("t_alert_msg").
		Select("t_alert_msg.id").
		Where("id = ? ", "").
		Find(&deptList).Error; err != nil {
	}
	return e
}

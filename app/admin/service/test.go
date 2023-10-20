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
	if err := e.Orm.Table("sys_role_dept").
		Select("sys_role_dept.dept_id").
		Joins("LEFT JOIN sys_dept on sys_dept.dept_id=sys_role_dept.dept_id").
		Where("role_id = ? ", "").
		Find(&deptList).Error; err != nil {
	}
	return e
}

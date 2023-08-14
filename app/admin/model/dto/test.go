package dto

type TestReq struct {
	Title    string `form:"title"  search:"type:contains;column:title;table:sys_api" comment:"标题"`
	Path     string `form:"path"  search:"type:contains;column:path;table:sys_api" comment:"地址"`
	Action   string `form:"action"  search:"type:exact;column:action;table:sys_api" comment:"请求方式"`
	ParentId string `form:"parentId"  search:"type:exact;column:parent_id;table:sys_api" comment:"按钮id"`
	Type     string `form:"type" search:"-" comment:"类型"`
}

func (s *TestReq) GetType() interface{} {
	return s.Type
}

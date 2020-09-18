package request

import (
	"anew-server/dto/response"
)

// User login structure
type RegisterAndLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 修改密码结构体
type ChangePwdReq struct {
	Req
	OldPassword string `json:"oldPassword" form:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" form:"newPassword" validate:"required"`
}

// 获取用户列表结构体
type UserListReq struct {
	Id       uint   `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Mobile   string `json:"mobile" form:"mobile"`
	Avatar   string `json:"avatar"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"mail" form:"mail"`
	Status   *bool  `json:"status" form:"status"`
	RoleId   uint   `json:"roleId" form:"roleId"`
	Creator  string `json:"creator" form:"creator"`
	response.PageInfo        // 分页参数
}

// 创建用户结构体
type CreateUserReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Mobile   string `json:"mobile" validate:"len=11"`
	Avatar   string `json:"avatar"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"mail"`
	Status   *bool  `json:"status"`
	RoleId   uint   `json:"roleId" validate:"required"`
	Creator  string `json:"creator"`
}

// 修改用户结构体
type UpdateUserReq struct {
	Username string `json:"username" validate:"required"`
	Mobile   string `json:"mobile" validate:"len=11"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"mail"`
	Password string `json:"password"`
	Status   *bool  `json:"status"`
}

// 翻译需要校验的字段名称
func (s CreateUserReq) FieldTrans() map[string]string {
	m := make(map[string]string, 0)
	m["Username"] = "用户名"
	m["Password"] = "用户密码"
	m["Mobile"] = "手机号码"
	m["Name"] = "姓名"
	m["RoleId"] = "角色ID"
	return m
}
func (s ChangePwdReq) FieldTrans() map[string]string {
	m := make(map[string]string, 0)
	m["OldPassword"] = "旧密码"
	m["NewPassword"] = "新密码"
	return m
}

func (s UpdateUserReq) FieldTrans() map[string]string {
	m := make(map[string]string, 0)
	m["Username"] = "用户名"
	m["Mobile"] = "手机号码"
	m["Name"] = "姓名"
	return m
}
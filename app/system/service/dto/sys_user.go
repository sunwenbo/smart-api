package dto

import (
	"smart-api/app/system/models"

	"smart-api/common/dto"
	common "smart-api/common/models"
)

type SysUserGetPageReq struct {
	dto.Pagination `search:"-"`
	UserId         int    `form:"userId" search:"type:exact;column:user_id;table:sys_user" comment:"用户ID"`
	Username       string `form:"username" search:"type:contains;column:username;table:sys_user" comment:"用户名"`
	NickName       string `form:"nickName" search:"type:contains;column:nick_name;table:sys_user" comment:"昵称"`
	Phone          string `form:"phone" search:"type:contains;column:phone;table:sys_user" comment:"手机号"`
	RoleId         string `form:"roleId" search:"type:exact;column:role_id;table:sys_user" comment:"角色ID"`
	Sex            string `form:"sex" search:"type:exact;column:sex;table:sys_user" comment:"性别"`
	Email          string `form:"email" search:"type:contains;column:email;table:sys_user" comment:"邮箱"`
	PostId         string `form:"postId" search:"type:exact;column:post_id;table:sys_user" comment:"岗位"`
	Status         string `form:"status" search:"type:exact;column:status;table:sys_user" comment:"状态"`
	DeptJoin       `search:"type:left;on:dept_id:dept_id;table:sys_user;join:sys_dept"`
	SysUserOrder
}
type SysUserGetSpecifyReq struct {
	dto.Pagination `search:"-"`
	UserId         int    `form:"userId" search:"type:exact;column:id;table:sys_user" comment:"用户ID"`
	Username       string `form:"username" search:"type:exact;column:username;table:sys_user" comment:"用户名"`
}

type SysUserOrder struct {
	UserIdOrder    string `search:"type:order;column:user_id;table:sys_user" form:"userIdOrder"`
	UsernameOrder  string `search:"type:order;column:username;table:sys_user" form:"usernameOrder"`
	StatusOrder    string `search:"type:order;column:status;table:sys_user" form:"statusOrder"`
	CreatedAtOrder string `search:"type:order;column:created_at;table:sys_user" form:"createdAtOrder"`
}

type DeptJoin struct {
	DeptId string `search:"type:contains;column:dept_path;table:sys_dept" form:"deptId"`
}

func (m *SysUserGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type ResetSysUserPwdReq struct {
	UserId   int    `json:"userId" comment:"用户ID" vd:"$>0"` // 用户ID
	Password string `json:"password" comment:"密码" vd:"len($)>0"`
	common.ControlBy
}

func (s *ResetSysUserPwdReq) GetId() interface{} {
	return s.UserId
}

func (s *ResetSysUserPwdReq) Generate(model *models.SysUser) {
	if s.UserId != 0 {
		model.UserId = s.UserId
	}
	model.Password = s.Password
}

type UpdateSysUserAvatarReq struct {
	UserId int    `json:"userId" comment:"用户ID" vd:"len($)>0"` // 用户ID
	Avatar string `json:"avatar" comment:"头像" vd:"len($)>0"`
	common.ControlBy
}

func (s *UpdateSysUserAvatarReq) GetId() interface{} {
	return s.UserId
}

func (s *UpdateSysUserAvatarReq) Generate(model *models.SysUser) {
	if s.UserId != 0 {
		model.UserId = s.UserId
	}
	model.Avatar = s.Avatar
}

type UpdateSysUserStatusReq struct {
	UserId int    `json:"userId" comment:"用户ID" vd:"$>0"` // 用户ID
	Status string `json:"status" comment:"状态" vd:"len($)>0"`
	common.ControlBy
}

func (s *UpdateSysUserStatusReq) GetId() interface{} {
	return s.UserId
}

func (s *UpdateSysUserStatusReq) Generate(model *models.SysUser) {
	if s.UserId != 0 {
		model.UserId = s.UserId
	}
	model.Status = s.Status
}

type SysUserInsertReq struct {
	UserId   int    `json:"userId" comment:"用户ID"` // 用户ID
	Username string `json:"username" comment:"用户名" vd:"len($)>0"`
	Password string `json:"password" comment:"密码"`
	NickName string `json:"nickName" comment:"昵称" vd:"len($)>0"`
	Phone    string `json:"phone" comment:"手机号" vd:"len($)>0"`
	RoleId   int    `json:"roleId" comment:"角色ID"`
	Avatar   string `json:"avatar" comment:"头像"`
	Sex      string `json:"sex" comment:"性别"`
	Email    string `json:"email" comment:"邮箱" vd:"len($)>0,email"`
	DeptId   int    `json:"deptId" comment:"部门" vd:"$>0"`
	PostId   int    `json:"postId" comment:"岗位"`
	Remark   string `json:"remark" comment:"备注"`
	Status   string `json:"status" comment:"状态" vd:"len($)>0" default:"1"`
	Source   string `json:"source" comment:"用户来源" vd:"len($)>0" default:"1"`
	common.ControlBy
}

type LdapUserInsertReq struct {
	Username string `json:"username" binding:"required" comment:"用户名" vd:"len($)>0"`
	Password string `json:"password" binding:"required" comment:"密码"`
	Source   string `json:"source" comment:"用户来源" vd:"len($)>0" default:"1"`
	Code     string `form:"Code" json:"code" binding:"required"`
	UUID     string `form:"UUID" json:"uuid" binding:"required"`
	RoleId   int    `json:"roleId" comment:"角色ID" default:"2"`
	Status   string `json:"status" comment:"状态"  default:"2"`
}

func (s *LdapUserInsertReq) Generate(model *models.SysUser) {
	model.Username = s.Username
	model.Password = s.Password
	model.Source = s.Source
	if s.RoleId == 0 {
		model.RoleId = 2 // 设置默认值
	}
	if s.Status == "" {
		model.Status = "2" // 设置默认值
	}
}

func (s *SysUserInsertReq) Generate(model *models.SysUser) {
	if s.UserId != 0 {
		model.UserId = s.UserId
	}
	model.Username = s.Username
	model.Password = s.Password
	model.NickName = s.NickName
	model.Phone = s.Phone
	model.RoleId = s.RoleId
	model.Avatar = s.Avatar
	model.Sex = s.Sex
	model.Email = s.Email
	model.DeptId = s.DeptId
	model.PostId = s.PostId
	model.Remark = s.Remark
	model.Status = s.Status
	model.Source = s.Source
	model.CreateBy = s.CreateBy
}

func (s *SysUserInsertReq) GetId() interface{} {
	return s.UserId
}

type SysUserUpdateReq struct {
	UserId   int    `json:"userId" comment:"用户ID"` // 用户ID
	Username string `json:"username" comment:"用户名" vd:"len($)>0"`
	NickName string `json:"nickName" comment:"昵称" vd:"len($)>0"`
	Phone    string `json:"phone" comment:"手机号" vd:"len($)>0"`
	RoleId   int    `json:"roleId" comment:"角色ID"`
	Avatar   string `json:"avatar" comment:"头像"`
	Sex      string `json:"sex" comment:"性别"`
	Email    string `json:"email" comment:"邮箱" vd:"len($)>0,email"`
	DeptId   int    `json:"deptId" comment:"部门" vd:"$>0"`
	PostId   int    `json:"postId" comment:"岗位"`
	Remark   string `json:"remark" comment:"备注"`
	Status   string `json:"status" comment:"状态" default:"1"`
	common.ControlBy
}

func (s *SysUserUpdateReq) Generate(model *models.SysUser) {
	if s.UserId != 0 {
		model.UserId = s.UserId
	}
	model.Username = s.Username
	model.NickName = s.NickName
	model.Phone = s.Phone
	model.RoleId = s.RoleId
	model.Avatar = s.Avatar
	model.Sex = s.Sex
	model.Email = s.Email
	model.DeptId = s.DeptId
	model.PostId = s.PostId
	model.Remark = s.Remark
	model.Status = s.Status
}

func (s *SysUserUpdateReq) GetId() interface{} {
	return s.UserId
}

type SysUserById struct {
	dto.ObjectById
	common.ControlBy
}

func (s *SysUserById) GetId() interface{} {
	if len(s.Ids) > 0 {
		s.Ids = append(s.Ids, s.Id)
		return s.Ids
	}
	return s.Id
}

func (s *SysUserById) GenerateM() (common.ActiveRecord, error) {
	return &models.SysUser{}, nil
}

// PassWord 密码
type PassWord struct {
	NewPassword string `json:"newPassword" vd:"len($)>0"`
	OldPassword string `json:"oldPassword" vd:"len($)>0"`
}

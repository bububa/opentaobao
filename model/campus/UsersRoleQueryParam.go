package campus

import (
	"sync"
)

// UsersRoleQueryParam 结构体
type UsersRoleQueryParam struct {
	// 多应用
	AppKeys []string `json:"app_keys,omitempty" xml:"app_keys>string,omitempty"`
	// 用户账号
	UserIds []string `json:"user_ids,omitempty" xml:"user_ids>string,omitempty"`
	// 角色key
	RoleKey string `json:"role_key,omitempty" xml:"role_key,omitempty"`
	// 角色名称
	RoleName string `json:"role_name,omitempty" xml:"role_name,omitempty"`
	// 园区id
	DeptId string `json:"dept_id,omitempty" xml:"dept_id,omitempty"`
	// 每页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 角色主键id
	RoleId int64 `json:"role_id,omitempty" xml:"role_id,omitempty"`
	// 当前页码
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
	// 是否支持有效期
	ReturnNotEffective bool `json:"return_not_effective,omitempty" xml:"return_not_effective,omitempty"`
	// true,返回用户拥有的所有角色；false 只返回roleId, roleName,roleType,roleKey过滤出来的角色
	ReturnAllUserRole bool `json:"return_all_user_role,omitempty" xml:"return_all_user_role,omitempty"`
}

var poolUsersRoleQueryParam = sync.Pool{
	New: func() any {
		return new(UsersRoleQueryParam)
	},
}

// GetUsersRoleQueryParam() 从对象池中获取UsersRoleQueryParam
func GetUsersRoleQueryParam() *UsersRoleQueryParam {
	return poolUsersRoleQueryParam.Get().(*UsersRoleQueryParam)
}

// ReleaseUsersRoleQueryParam 释放UsersRoleQueryParam
func ReleaseUsersRoleQueryParam(v *UsersRoleQueryParam) {
	v.AppKeys = v.AppKeys[:0]
	v.UserIds = v.UserIds[:0]
	v.RoleKey = ""
	v.RoleName = ""
	v.DeptId = ""
	v.PageSize = 0
	v.RoleId = 0
	v.PageNum = 0
	v.ReturnNotEffective = false
	v.ReturnAllUserRole = false
	poolUsersRoleQueryParam.Put(v)
}

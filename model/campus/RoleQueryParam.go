package campus

import (
	"sync"
)

// RoleQueryParam 结构体
type RoleQueryParam struct {
	// 支持多应用
	AppKeys []string `json:"app_keys,omitempty" xml:"app_keys>string,omitempty"`
	// 角色名称
	RoleName string `json:"role_name,omitempty" xml:"role_name,omitempty"`
	// 园区
	DeptId string `json:"dept_id,omitempty" xml:"dept_id,omitempty"`
}

var poolRoleQueryParam = sync.Pool{
	New: func() any {
		return new(RoleQueryParam)
	},
}

// GetRoleQueryParam() 从对象池中获取RoleQueryParam
func GetRoleQueryParam() *RoleQueryParam {
	return poolRoleQueryParam.Get().(*RoleQueryParam)
}

// ReleaseRoleQueryParam 释放RoleQueryParam
func ReleaseRoleQueryParam(v *RoleQueryParam) {
	v.AppKeys = v.AppKeys[:0]
	v.RoleName = ""
	v.DeptId = ""
	poolRoleQueryParam.Put(v)
}

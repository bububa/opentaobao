package campus

import (
	"sync"
)

// SysRolePermissionsVo 结构体
type SysRolePermissionsVo struct {
	// permissions
	Permissions []TreeNodeDto `json:"permissions,omitempty" xml:"permissions>tree_node_dto,omitempty"`
	// 角色key
	RoleKey string `json:"role_key,omitempty" xml:"role_key,omitempty"`
	// 租户
	Tenant string `json:"tenant,omitempty" xml:"tenant,omitempty"`
	// appKey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 1
	DeptId string `json:"dept_id,omitempty" xml:"dept_id,omitempty"`
	// 园区名称
	DeptName string `json:"dept_name,omitempty" xml:"dept_name,omitempty"`
	// 角色名称
	RoleName string `json:"role_name,omitempty" xml:"role_name,omitempty"`
	// 角色类型
	RoleType string `json:"role_type,omitempty" xml:"role_type,omitempty"`
	// 角色描述
	RoleDesc string `json:"role_desc,omitempty" xml:"role_desc,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolSysRolePermissionsVo = sync.Pool{
	New: func() any {
		return new(SysRolePermissionsVo)
	},
}

// GetSysRolePermissionsVo() 从对象池中获取SysRolePermissionsVo
func GetSysRolePermissionsVo() *SysRolePermissionsVo {
	return poolSysRolePermissionsVo.Get().(*SysRolePermissionsVo)
}

// ReleaseSysRolePermissionsVo 释放SysRolePermissionsVo
func ReleaseSysRolePermissionsVo(v *SysRolePermissionsVo) {
	v.Permissions = v.Permissions[:0]
	v.RoleKey = ""
	v.Tenant = ""
	v.AppKey = ""
	v.DeptId = ""
	v.DeptName = ""
	v.RoleName = ""
	v.RoleType = ""
	v.RoleDesc = ""
	v.Id = 0
	poolSysRolePermissionsVo.Put(v)
}

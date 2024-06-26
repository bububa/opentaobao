package campus

import (
	"sync"
)

// SysRoleDto 结构体
type SysRoleDto struct {
	// 角色key
	RoleKey string `json:"role_key,omitempty" xml:"role_key,omitempty"`
	// 角色属于哪个租户
	Tenant string `json:"tenant,omitempty" xml:"tenant,omitempty"`
	// 应用
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 园区
	DeptId string `json:"dept_id,omitempty" xml:"dept_id,omitempty"`
	// 园区名称
	DeptName string `json:"dept_name,omitempty" xml:"dept_name,omitempty"`
	// 角色名称
	RoleName string `json:"role_name,omitempty" xml:"role_name,omitempty"`
	// 角色描述
	RoleDesc string `json:"role_desc,omitempty" xml:"role_desc,omitempty"`
	// 角色类型
	RoleType string `json:"role_type,omitempty" xml:"role_type,omitempty"`
	// 修改人id
	ModifierId string `json:"modifier_id,omitempty" xml:"modifier_id,omitempty"`
	// 创建人id
	CreatorId string `json:"creator_id,omitempty" xml:"creator_id,omitempty"`
	// 角色主键id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolSysRoleDto = sync.Pool{
	New: func() any {
		return new(SysRoleDto)
	},
}

// GetSysRoleDto() 从对象池中获取SysRoleDto
func GetSysRoleDto() *SysRoleDto {
	return poolSysRoleDto.Get().(*SysRoleDto)
}

// ReleaseSysRoleDto 释放SysRoleDto
func ReleaseSysRoleDto(v *SysRoleDto) {
	v.RoleKey = ""
	v.Tenant = ""
	v.AppKey = ""
	v.DeptId = ""
	v.DeptName = ""
	v.RoleName = ""
	v.RoleDesc = ""
	v.RoleType = ""
	v.ModifierId = ""
	v.CreatorId = ""
	v.Id = 0
	poolSysRoleDto.Put(v)
}

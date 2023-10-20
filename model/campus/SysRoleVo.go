package campus

import (
	"sync"
)

// SysRoleVo 结构体
type SysRoleVo struct {
	// id
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// roleName
	RoleName string `json:"role_name,omitempty" xml:"role_name,omitempty"`
	// 是否被用户选择
	Granted string `json:"granted,omitempty" xml:"granted,omitempty"`
	// 园区名称
	DeptName string `json:"dept_name,omitempty" xml:"dept_name,omitempty"`
}

var poolSysRoleVo = sync.Pool{
	New: func() any {
		return new(SysRoleVo)
	},
}

// GetSysRoleVo() 从对象池中获取SysRoleVo
func GetSysRoleVo() *SysRoleVo {
	return poolSysRoleVo.Get().(*SysRoleVo)
}

// ReleaseSysRoleVo 释放SysRoleVo
func ReleaseSysRoleVo(v *SysRoleVo) {
	v.Id = ""
	v.RoleName = ""
	v.Granted = ""
	v.DeptName = ""
	poolSysRoleVo.Put(v)
}

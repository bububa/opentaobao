package subuser

import (
	"sync"
)

// Permission 结构体
type Permission struct {
	// 注册到权限中心的code值
	PermissionCode string `json:"permission_code,omitempty" xml:"permission_code,omitempty"`
	// 权限名称
	PermissionName string `json:"permission_name,omitempty" xml:"permission_name,omitempty"`
	// 父权限code
	ParentCode string `json:"parent_code,omitempty" xml:"parent_code,omitempty"`
	// 1 :允许授权  2：不允许授权 6：不允许授权但默认已有权限
	IsAuthorize int64 `json:"is_authorize,omitempty" xml:"is_authorize,omitempty"`
}

var poolPermission = sync.Pool{
	New: func() any {
		return new(Permission)
	},
}

// GetPermission() 从对象池中获取Permission
func GetPermission() *Permission {
	return poolPermission.Get().(*Permission)
}

// ReleasePermission 释放Permission
func ReleasePermission(v *Permission) {
	v.PermissionCode = ""
	v.PermissionName = ""
	v.ParentCode = ""
	v.IsAuthorize = 0
	poolPermission.Put(v)
}

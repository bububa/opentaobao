package campus

import (
	"sync"
)

// CheckUserPermissionParam 结构体
type CheckUserPermissionParam struct {
	// 权限key
	PermissionKey string `json:"permission_key,omitempty" xml:"permission_key,omitempty"`
	// 用户账号
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 是否关联校验权限组下的权限
	HierarchicalCheckPermissionGroup bool `json:"hierarchical_check_permission_group,omitempty" xml:"hierarchical_check_permission_group,omitempty"`
	// 是否关联角色和权限组下的权限校验
	HierarchicalCheckRolePermissionGroup bool `json:"hierarchical_check_role_permission_group,omitempty" xml:"hierarchical_check_role_permission_group,omitempty"`
	// 是否用户拥有菜单，就认为用户级联拥有菜单下的权限
	HierarchicalObtainMenuPermission bool `json:"hierarchical_obtain_menu_permission,omitempty" xml:"hierarchical_obtain_menu_permission,omitempty"`
}

var poolCheckUserPermissionParam = sync.Pool{
	New: func() any {
		return new(CheckUserPermissionParam)
	},
}

// GetCheckUserPermissionParam() 从对象池中获取CheckUserPermissionParam
func GetCheckUserPermissionParam() *CheckUserPermissionParam {
	return poolCheckUserPermissionParam.Get().(*CheckUserPermissionParam)
}

// ReleaseCheckUserPermissionParam 释放CheckUserPermissionParam
func ReleaseCheckUserPermissionParam(v *CheckUserPermissionParam) {
	v.PermissionKey = ""
	v.UserId = ""
	v.HierarchicalCheckPermissionGroup = false
	v.HierarchicalCheckRolePermissionGroup = false
	v.HierarchicalObtainMenuPermission = false
	poolCheckUserPermissionParam.Put(v)
}

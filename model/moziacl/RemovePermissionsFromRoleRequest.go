package moziacl

import (
	"sync"
)

// RemovePermissionsFromRoleRequest 结构体
type RemovePermissionsFromRoleRequest struct {
	// 要移除的权限name列表(功能权限唯一code，在ACL中全局唯一)
	PermissionNames []string `json:"permission_names,omitempty" xml:"permission_names>string,omitempty"`
	// 角色所在应用name(应用唯一code)
	TargetAppName string `json:"target_app_name,omitempty" xml:"target_app_name,omitempty"`
	// 角色name(角色唯一code，在ACL中全局唯一)
	RoleName string `json:"role_name,omitempty" xml:"role_name,omitempty"`
	// 请求参数扩展字段
	RequestMetaData string `json:"request_meta_data,omitempty" xml:"request_meta_data,omitempty"`
	// 操作主体
	PrincipalParam *BucUserPrincipalParam `json:"principal_param,omitempty" xml:"principal_param,omitempty"`
}

var poolRemovePermissionsFromRoleRequest = sync.Pool{
	New: func() any {
		return new(RemovePermissionsFromRoleRequest)
	},
}

// GetRemovePermissionsFromRoleRequest() 从对象池中获取RemovePermissionsFromRoleRequest
func GetRemovePermissionsFromRoleRequest() *RemovePermissionsFromRoleRequest {
	return poolRemovePermissionsFromRoleRequest.Get().(*RemovePermissionsFromRoleRequest)
}

// ReleaseRemovePermissionsFromRoleRequest 释放RemovePermissionsFromRoleRequest
func ReleaseRemovePermissionsFromRoleRequest(v *RemovePermissionsFromRoleRequest) {
	v.PermissionNames = v.PermissionNames[:0]
	v.TargetAppName = ""
	v.RoleName = ""
	v.RequestMetaData = ""
	v.PrincipalParam = nil
	poolRemovePermissionsFromRoleRequest.Put(v)
}

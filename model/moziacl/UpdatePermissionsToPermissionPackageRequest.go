package moziacl

import (
	"sync"
)

// UpdatePermissionsToPermissionPackageRequest 结构体
type UpdatePermissionsToPermissionPackageRequest struct {
	// 权限唯一标识列表
	PermissionNames []string `json:"permission_names,omitempty" xml:"permission_names>string,omitempty"`
	// 权限套餐唯一标识
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 请求扩展字段，返回数据中扩展字段值与此相同
	RequestMetaData string `json:"request_meta_data,omitempty" xml:"request_meta_data,omitempty"`
	// 请求主体
	PrincipalParam *BucUserPrincipalParam `json:"principal_param,omitempty" xml:"principal_param,omitempty"`
}

var poolUpdatePermissionsToPermissionPackageRequest = sync.Pool{
	New: func() any {
		return new(UpdatePermissionsToPermissionPackageRequest)
	},
}

// GetUpdatePermissionsToPermissionPackageRequest() 从对象池中获取UpdatePermissionsToPermissionPackageRequest
func GetUpdatePermissionsToPermissionPackageRequest() *UpdatePermissionsToPermissionPackageRequest {
	return poolUpdatePermissionsToPermissionPackageRequest.Get().(*UpdatePermissionsToPermissionPackageRequest)
}

// ReleaseUpdatePermissionsToPermissionPackageRequest 释放UpdatePermissionsToPermissionPackageRequest
func ReleaseUpdatePermissionsToPermissionPackageRequest(v *UpdatePermissionsToPermissionPackageRequest) {
	v.PermissionNames = v.PermissionNames[:0]
	v.Name = ""
	v.RequestMetaData = ""
	v.PrincipalParam = nil
	poolUpdatePermissionsToPermissionPackageRequest.Put(v)
}

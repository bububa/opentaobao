package moziacl

import (
	"sync"
)

// UpdatePermissionsToPermissionPackageResult 结构体
type UpdatePermissionsToPermissionPackageResult struct {
	// 返回data，此接口此字段返回值为空
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 请求唯一id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应message
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// 扩展字段，值与入参扩展字段相同
	ResponseMetaData string `json:"response_meta_data,omitempty" xml:"response_meta_data,omitempty"`
	// 响应code
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolUpdatePermissionsToPermissionPackageResult = sync.Pool{
	New: func() any {
		return new(UpdatePermissionsToPermissionPackageResult)
	},
}

// GetUpdatePermissionsToPermissionPackageResult() 从对象池中获取UpdatePermissionsToPermissionPackageResult
func GetUpdatePermissionsToPermissionPackageResult() *UpdatePermissionsToPermissionPackageResult {
	return poolUpdatePermissionsToPermissionPackageResult.Get().(*UpdatePermissionsToPermissionPackageResult)
}

// ReleaseUpdatePermissionsToPermissionPackageResult 释放UpdatePermissionsToPermissionPackageResult
func ReleaseUpdatePermissionsToPermissionPackageResult(v *UpdatePermissionsToPermissionPackageResult) {
	v.Data = ""
	v.RequestId = ""
	v.ResponseMessage = ""
	v.ResponseMetaData = ""
	v.ResponseCode = ""
	v.Success = false
	poolUpdatePermissionsToPermissionPackageResult.Put(v)
}

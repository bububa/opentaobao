package moziacl

import (
	"sync"
)

// AddPermissionToRoleResult 结构体
type AddPermissionToRoleResult struct {
	// 返回数据，这个接口返回数据为空
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 请求唯一id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果message，若处理失败则返回失败原因
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// 扩展字段，与入参扩展字段值相同
	ResponseMetaData string `json:"response_meta_data,omitempty" xml:"response_meta_data,omitempty"`
	// 响应code
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 是否处理成功，若成功则返回true
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAddPermissionToRoleResult = sync.Pool{
	New: func() any {
		return new(AddPermissionToRoleResult)
	},
}

// GetAddPermissionToRoleResult() 从对象池中获取AddPermissionToRoleResult
func GetAddPermissionToRoleResult() *AddPermissionToRoleResult {
	return poolAddPermissionToRoleResult.Get().(*AddPermissionToRoleResult)
}

// ReleaseAddPermissionToRoleResult 释放AddPermissionToRoleResult
func ReleaseAddPermissionToRoleResult(v *AddPermissionToRoleResult) {
	v.Data = ""
	v.RequestId = ""
	v.ResponseMessage = ""
	v.ResponseMetaData = ""
	v.ResponseCode = ""
	v.Success = false
	poolAddPermissionToRoleResult.Put(v)
}

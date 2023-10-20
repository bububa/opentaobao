package moziacl

import (
	"sync"
)

// RevokePermissionsResult 结构体
type RevokePermissionsResult struct {
	// 请求唯一id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应message，若失败则返回失败原因
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// 扩展字段，与入参扩展字段值相同
	ResponseMetaData string `json:"response_meta_data,omitempty" xml:"response_meta_data,omitempty"`
	// 响应code
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 是否调用成功，成功则为true
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolRevokePermissionsResult = sync.Pool{
	New: func() any {
		return new(RevokePermissionsResult)
	},
}

// GetRevokePermissionsResult() 从对象池中获取RevokePermissionsResult
func GetRevokePermissionsResult() *RevokePermissionsResult {
	return poolRevokePermissionsResult.Get().(*RevokePermissionsResult)
}

// ReleaseRevokePermissionsResult 释放RevokePermissionsResult
func ReleaseRevokePermissionsResult(v *RevokePermissionsResult) {
	v.RequestId = ""
	v.ResponseMessage = ""
	v.ResponseMetaData = ""
	v.ResponseCode = ""
	v.Success = false
	poolRevokePermissionsResult.Put(v)
}

package mozivds

import (
	"sync"
)

// RemoveTenantAdminsResult 结构体
type RemoveTenantAdminsResult struct {
	// 请求id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回消息
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// 返回元数据
	ResponseMetaData string `json:"response_meta_data,omitempty" xml:"response_meta_data,omitempty"`
	// 返回Code
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolRemoveTenantAdminsResult = sync.Pool{
	New: func() any {
		return new(RemoveTenantAdminsResult)
	},
}

// GetRemoveTenantAdminsResult() 从对象池中获取RemoveTenantAdminsResult
func GetRemoveTenantAdminsResult() *RemoveTenantAdminsResult {
	return poolRemoveTenantAdminsResult.Get().(*RemoveTenantAdminsResult)
}

// ReleaseRemoveTenantAdminsResult 释放RemoveTenantAdminsResult
func ReleaseRemoveTenantAdminsResult(v *RemoveTenantAdminsResult) {
	v.RequestId = ""
	v.ResponseMessage = ""
	v.ResponseMetaData = ""
	v.ResponseCode = ""
	v.Success = false
	poolRemoveTenantAdminsResult.Put(v)
}

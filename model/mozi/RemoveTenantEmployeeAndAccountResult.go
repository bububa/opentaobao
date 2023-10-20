package mozi

import (
	"sync"
)

// RemoveTenantEmployeeAndAccountResult 结构体
type RemoveTenantEmployeeAndAccountResult struct {
	// 返回状态描述
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// 返回附加信息
	ResponseMetaData string `json:"response_meta_data,omitempty" xml:"response_meta_data,omitempty"`
	// 返回的状态码
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 请求的UUID
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolRemoveTenantEmployeeAndAccountResult = sync.Pool{
	New: func() any {
		return new(RemoveTenantEmployeeAndAccountResult)
	},
}

// GetRemoveTenantEmployeeAndAccountResult() 从对象池中获取RemoveTenantEmployeeAndAccountResult
func GetRemoveTenantEmployeeAndAccountResult() *RemoveTenantEmployeeAndAccountResult {
	return poolRemoveTenantEmployeeAndAccountResult.Get().(*RemoveTenantEmployeeAndAccountResult)
}

// ReleaseRemoveTenantEmployeeAndAccountResult 释放RemoveTenantEmployeeAndAccountResult
func ReleaseRemoveTenantEmployeeAndAccountResult(v *RemoveTenantEmployeeAndAccountResult) {
	v.ResponseMessage = ""
	v.ResponseMetaData = ""
	v.ResponseCode = ""
	v.RequestId = ""
	v.Success = false
	poolRemoveTenantEmployeeAndAccountResult.Put(v)
}

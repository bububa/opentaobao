package mozi

import (
	"sync"
)

// UpdateTenantEmployeeAndAccountResult 结构体
type UpdateTenantEmployeeAndAccountResult struct {
	// 返回结果成功还是失败
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 请求的UUID
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回状态描述
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// 返回附加信息
	ResponseMetaData string `json:"response_meta_data,omitempty" xml:"response_meta_data,omitempty"`
	// 返回的状态码
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolUpdateTenantEmployeeAndAccountResult = sync.Pool{
	New: func() any {
		return new(UpdateTenantEmployeeAndAccountResult)
	},
}

// GetUpdateTenantEmployeeAndAccountResult() 从对象池中获取UpdateTenantEmployeeAndAccountResult
func GetUpdateTenantEmployeeAndAccountResult() *UpdateTenantEmployeeAndAccountResult {
	return poolUpdateTenantEmployeeAndAccountResult.Get().(*UpdateTenantEmployeeAndAccountResult)
}

// ReleaseUpdateTenantEmployeeAndAccountResult 释放UpdateTenantEmployeeAndAccountResult
func ReleaseUpdateTenantEmployeeAndAccountResult(v *UpdateTenantEmployeeAndAccountResult) {
	v.Data = ""
	v.RequestId = ""
	v.ResponseMessage = ""
	v.ResponseMetaData = ""
	v.ResponseCode = ""
	v.Success = false
	poolUpdateTenantEmployeeAndAccountResult.Put(v)
}

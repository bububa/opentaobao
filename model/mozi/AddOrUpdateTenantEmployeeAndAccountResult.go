package mozi

import (
	"sync"
)

// AddOrUpdateTenantEmployeeAndAccountResult 结构体
type AddOrUpdateTenantEmployeeAndAccountResult struct {
	// 返回状态描述
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// 返回附加信息
	ResponseMetaData string `json:"response_meta_data,omitempty" xml:"response_meta_data,omitempty"`
	// 返回的状态码
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 请求的UUID
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的人员和账号的绑定对象
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAddOrUpdateTenantEmployeeAndAccountResult = sync.Pool{
	New: func() any {
		return new(AddOrUpdateTenantEmployeeAndAccountResult)
	},
}

// GetAddOrUpdateTenantEmployeeAndAccountResult() 从对象池中获取AddOrUpdateTenantEmployeeAndAccountResult
func GetAddOrUpdateTenantEmployeeAndAccountResult() *AddOrUpdateTenantEmployeeAndAccountResult {
	return poolAddOrUpdateTenantEmployeeAndAccountResult.Get().(*AddOrUpdateTenantEmployeeAndAccountResult)
}

// ReleaseAddOrUpdateTenantEmployeeAndAccountResult 释放AddOrUpdateTenantEmployeeAndAccountResult
func ReleaseAddOrUpdateTenantEmployeeAndAccountResult(v *AddOrUpdateTenantEmployeeAndAccountResult) {
	v.ResponseMessage = ""
	v.ResponseMetaData = ""
	v.ResponseCode = ""
	v.RequestId = ""
	v.Data = ""
	v.Success = false
	poolAddOrUpdateTenantEmployeeAndAccountResult.Put(v)
}

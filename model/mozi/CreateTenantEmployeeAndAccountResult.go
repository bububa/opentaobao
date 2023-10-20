package mozi

import (
	"sync"
)

// CreateTenantEmployeeAndAccountResult 结构体
type CreateTenantEmployeeAndAccountResult struct {
	// 返回状态描述
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// 返回附加信息
	ResponseMetaData string `json:"response_meta_data,omitempty" xml:"response_meta_data,omitempty"`
	// 返回的状态码
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 请求的UUID
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的人员和账号的绑定对象
	Data *EmployeeAccount `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCreateTenantEmployeeAndAccountResult = sync.Pool{
	New: func() any {
		return new(CreateTenantEmployeeAndAccountResult)
	},
}

// GetCreateTenantEmployeeAndAccountResult() 从对象池中获取CreateTenantEmployeeAndAccountResult
func GetCreateTenantEmployeeAndAccountResult() *CreateTenantEmployeeAndAccountResult {
	return poolCreateTenantEmployeeAndAccountResult.Get().(*CreateTenantEmployeeAndAccountResult)
}

// ReleaseCreateTenantEmployeeAndAccountResult 释放CreateTenantEmployeeAndAccountResult
func ReleaseCreateTenantEmployeeAndAccountResult(v *CreateTenantEmployeeAndAccountResult) {
	v.ResponseMessage = ""
	v.ResponseMetaData = ""
	v.ResponseCode = ""
	v.RequestId = ""
	v.Data = nil
	v.Success = false
	poolCreateTenantEmployeeAndAccountResult.Put(v)
}

package mozi

import (
	"sync"
)

// GetTenantByIdResult 结构体
type GetTenantByIdResult struct {
	// requestId
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// responseMessage
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// responseCode
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// data数据
	Data *Tenant `json:"data,omitempty" xml:"data,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolGetTenantByIdResult = sync.Pool{
	New: func() any {
		return new(GetTenantByIdResult)
	},
}

// GetGetTenantByIdResult() 从对象池中获取GetTenantByIdResult
func GetGetTenantByIdResult() *GetTenantByIdResult {
	return poolGetTenantByIdResult.Get().(*GetTenantByIdResult)
}

// ReleaseGetTenantByIdResult 释放GetTenantByIdResult
func ReleaseGetTenantByIdResult(v *GetTenantByIdResult) {
	v.RequestId = ""
	v.ResponseMessage = ""
	v.ResponseCode = ""
	v.Data = nil
	v.Success = false
	poolGetTenantByIdResult.Put(v)
}

package mozi

import (
	"sync"
)

// GetEmployeeTenantAdminInfoResult 结构体
type GetEmployeeTenantAdminInfoResult struct {
	// requestId
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// responseMessage
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// responseCode
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否是管理员
	Admin bool `json:"admin,omitempty" xml:"admin,omitempty"`
	// 是否主管理员
	Primary bool `json:"primary,omitempty" xml:"primary,omitempty"`
}

var poolGetEmployeeTenantAdminInfoResult = sync.Pool{
	New: func() any {
		return new(GetEmployeeTenantAdminInfoResult)
	},
}

// GetGetEmployeeTenantAdminInfoResult() 从对象池中获取GetEmployeeTenantAdminInfoResult
func GetGetEmployeeTenantAdminInfoResult() *GetEmployeeTenantAdminInfoResult {
	return poolGetEmployeeTenantAdminInfoResult.Get().(*GetEmployeeTenantAdminInfoResult)
}

// ReleaseGetEmployeeTenantAdminInfoResult 释放GetEmployeeTenantAdminInfoResult
func ReleaseGetEmployeeTenantAdminInfoResult(v *GetEmployeeTenantAdminInfoResult) {
	v.RequestId = ""
	v.ResponseMessage = ""
	v.ResponseCode = ""
	v.Success = false
	v.Admin = false
	v.Primary = false
	poolGetEmployeeTenantAdminInfoResult.Put(v)
}

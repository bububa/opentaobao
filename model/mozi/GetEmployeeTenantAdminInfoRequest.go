package mozi

import (
	"sync"
)

// GetEmployeeTenantAdminInfoRequest 结构体
type GetEmployeeTenantAdminInfoRequest struct {
	// 人员code
	EmployeeCode string `json:"employee_code,omitempty" xml:"employee_code,omitempty"`
	// 租户ID
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
}

var poolGetEmployeeTenantAdminInfoRequest = sync.Pool{
	New: func() any {
		return new(GetEmployeeTenantAdminInfoRequest)
	},
}

// GetGetEmployeeTenantAdminInfoRequest() 从对象池中获取GetEmployeeTenantAdminInfoRequest
func GetGetEmployeeTenantAdminInfoRequest() *GetEmployeeTenantAdminInfoRequest {
	return poolGetEmployeeTenantAdminInfoRequest.Get().(*GetEmployeeTenantAdminInfoRequest)
}

// ReleaseGetEmployeeTenantAdminInfoRequest 释放GetEmployeeTenantAdminInfoRequest
func ReleaseGetEmployeeTenantAdminInfoRequest(v *GetEmployeeTenantAdminInfoRequest) {
	v.EmployeeCode = ""
	v.TenantId = 0
	poolGetEmployeeTenantAdminInfoRequest.Put(v)
}

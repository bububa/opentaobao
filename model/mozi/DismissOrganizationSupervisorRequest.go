package mozi

import (
	"sync"
)

// DismissOrganizationSupervisorRequest 结构体
type DismissOrganizationSupervisorRequest struct {
	// 员工 Code 列表
	EmployeeCodes []string `json:"employee_codes,omitempty" xml:"employee_codes>string,omitempty"`
	// 组织 Code
	OrganizationCode string `json:"organization_code,omitempty" xml:"organization_code,omitempty"`
	// 操作人
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 租户 ID
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
}

var poolDismissOrganizationSupervisorRequest = sync.Pool{
	New: func() any {
		return new(DismissOrganizationSupervisorRequest)
	},
}

// GetDismissOrganizationSupervisorRequest() 从对象池中获取DismissOrganizationSupervisorRequest
func GetDismissOrganizationSupervisorRequest() *DismissOrganizationSupervisorRequest {
	return poolDismissOrganizationSupervisorRequest.Get().(*DismissOrganizationSupervisorRequest)
}

// ReleaseDismissOrganizationSupervisorRequest 释放DismissOrganizationSupervisorRequest
func ReleaseDismissOrganizationSupervisorRequest(v *DismissOrganizationSupervisorRequest) {
	v.EmployeeCodes = v.EmployeeCodes[:0]
	v.OrganizationCode = ""
	v.Operator = ""
	v.TenantId = 0
	poolDismissOrganizationSupervisorRequest.Put(v)
}

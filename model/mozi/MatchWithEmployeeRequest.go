package mozi

import (
	"sync"
)

// MatchWithEmployeeRequest 结构体
type MatchWithEmployeeRequest struct {
	// 人员code
	EmployeeCodes []string `json:"employee_codes,omitempty" xml:"employee_codes>string,omitempty"`
	// 组code
	GroupCode string `json:"group_code,omitempty" xml:"group_code,omitempty"`
	// 租户id
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
}

var poolMatchWithEmployeeRequest = sync.Pool{
	New: func() any {
		return new(MatchWithEmployeeRequest)
	},
}

// GetMatchWithEmployeeRequest() 从对象池中获取MatchWithEmployeeRequest
func GetMatchWithEmployeeRequest() *MatchWithEmployeeRequest {
	return poolMatchWithEmployeeRequest.Get().(*MatchWithEmployeeRequest)
}

// ReleaseMatchWithEmployeeRequest 释放MatchWithEmployeeRequest
func ReleaseMatchWithEmployeeRequest(v *MatchWithEmployeeRequest) {
	v.EmployeeCodes = v.EmployeeCodes[:0]
	v.GroupCode = ""
	v.TenantId = 0
	poolMatchWithEmployeeRequest.Put(v)
}

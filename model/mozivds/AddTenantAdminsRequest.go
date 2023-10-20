package mozivds

import (
	"sync"
)

// AddTenantAdminsRequest 结构体
type AddTenantAdminsRequest struct {
	// 人员code
	EmployeeCodes []string `json:"employee_codes,omitempty" xml:"employee_codes>string,omitempty"`
	// 操作人
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 请求元数据
	RequestMetaData string `json:"request_meta_data,omitempty" xml:"request_meta_data,omitempty"`
	// 租户Id
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
	// 是否主管理员
	PrimaryAdmin bool `json:"primary_admin,omitempty" xml:"primary_admin,omitempty"`
}

var poolAddTenantAdminsRequest = sync.Pool{
	New: func() any {
		return new(AddTenantAdminsRequest)
	},
}

// GetAddTenantAdminsRequest() 从对象池中获取AddTenantAdminsRequest
func GetAddTenantAdminsRequest() *AddTenantAdminsRequest {
	return poolAddTenantAdminsRequest.Get().(*AddTenantAdminsRequest)
}

// ReleaseAddTenantAdminsRequest 释放AddTenantAdminsRequest
func ReleaseAddTenantAdminsRequest(v *AddTenantAdminsRequest) {
	v.EmployeeCodes = v.EmployeeCodes[:0]
	v.Operator = ""
	v.RequestMetaData = ""
	v.TenantId = 0
	v.PrimaryAdmin = false
	poolAddTenantAdminsRequest.Put(v)
}

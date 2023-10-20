package mozi

import (
	"sync"
)

// UpdateTenantEmployeeAndAccountRequest 结构体
type UpdateTenantEmployeeAndAccountRequest struct {
	// 员工姓名
	EmployeeName string `json:"employee_name,omitempty" xml:"employee_name,omitempty"`
	// 证件号码
	CertificateCode string `json:"certificate_code,omitempty" xml:"certificate_code,omitempty"`
	// 操作人
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 请求附加消息
	RequestMetaData string `json:"request_meta_data,omitempty" xml:"request_meta_data,omitempty"`
	// 员工CODE
	EmployeeCode string `json:"employee_code,omitempty" xml:"employee_code,omitempty"`
	// 证件类型
	CertificateType string `json:"certificate_type,omitempty" xml:"certificate_type,omitempty"`
	// 员工基本属性
	EmployeeBaseProperties *EmployeeBaseProperties `json:"employee_base_properties,omitempty" xml:"employee_base_properties,omitempty"`
	// 账号ID
	AccountId int64 `json:"account_id,omitempty" xml:"account_id,omitempty"`
	// 租户ID
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
}

var poolUpdateTenantEmployeeAndAccountRequest = sync.Pool{
	New: func() any {
		return new(UpdateTenantEmployeeAndAccountRequest)
	},
}

// GetUpdateTenantEmployeeAndAccountRequest() 从对象池中获取UpdateTenantEmployeeAndAccountRequest
func GetUpdateTenantEmployeeAndAccountRequest() *UpdateTenantEmployeeAndAccountRequest {
	return poolUpdateTenantEmployeeAndAccountRequest.Get().(*UpdateTenantEmployeeAndAccountRequest)
}

// ReleaseUpdateTenantEmployeeAndAccountRequest 释放UpdateTenantEmployeeAndAccountRequest
func ReleaseUpdateTenantEmployeeAndAccountRequest(v *UpdateTenantEmployeeAndAccountRequest) {
	v.EmployeeName = ""
	v.CertificateCode = ""
	v.Operator = ""
	v.RequestMetaData = ""
	v.EmployeeCode = ""
	v.CertificateType = ""
	v.EmployeeBaseProperties = nil
	v.AccountId = 0
	v.TenantId = 0
	poolUpdateTenantEmployeeAndAccountRequest.Put(v)
}

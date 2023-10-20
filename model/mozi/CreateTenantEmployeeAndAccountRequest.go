package mozi

import (
	"sync"
)

// CreateTenantEmployeeAndAccountRequest 结构体
type CreateTenantEmployeeAndAccountRequest struct {
	// 证件号码
	CertificateCode string `json:"certificate_code,omitempty" xml:"certificate_code,omitempty"`
	// 请求附加消息
	RequestMetaData string `json:"request_meta_data,omitempty" xml:"request_meta_data,omitempty"`
	// 昵称
	Nickname string `json:"nickname,omitempty" xml:"nickname,omitempty"`
	// 证件类型
	CertificateType string `json:"certificate_type,omitempty" xml:"certificate_type,omitempty"`
	// 员工CODE
	EmployeeCode string `json:"employee_code,omitempty" xml:"employee_code,omitempty"`
	// 密码
	Password string `json:"password,omitempty" xml:"password,omitempty"`
	// 操作人
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 员工姓名
	EmployeeName string `json:"employee_name,omitempty" xml:"employee_name,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 账户名
	Account string `json:"account,omitempty" xml:"account,omitempty"`
	// 语言
	Language string `json:"language,omitempty" xml:"language,omitempty"`
	// 手机区号
	SecMobileAreaCode string `json:"sec_mobile_area_code,omitempty" xml:"sec_mobile_area_code,omitempty"`
	// 工号
	EmployeeNumber string `json:"employee_number,omitempty" xml:"employee_number,omitempty"`
	// 手机号码
	SecMobile string `json:"sec_mobile,omitempty" xml:"sec_mobile,omitempty"`
	// 自建账号定义的namespace
	Namespace string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// 头像的完整URL
	Avatar string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// 租户ID
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
}

var poolCreateTenantEmployeeAndAccountRequest = sync.Pool{
	New: func() any {
		return new(CreateTenantEmployeeAndAccountRequest)
	},
}

// GetCreateTenantEmployeeAndAccountRequest() 从对象池中获取CreateTenantEmployeeAndAccountRequest
func GetCreateTenantEmployeeAndAccountRequest() *CreateTenantEmployeeAndAccountRequest {
	return poolCreateTenantEmployeeAndAccountRequest.Get().(*CreateTenantEmployeeAndAccountRequest)
}

// ReleaseCreateTenantEmployeeAndAccountRequest 释放CreateTenantEmployeeAndAccountRequest
func ReleaseCreateTenantEmployeeAndAccountRequest(v *CreateTenantEmployeeAndAccountRequest) {
	v.CertificateCode = ""
	v.RequestMetaData = ""
	v.Nickname = ""
	v.CertificateType = ""
	v.EmployeeCode = ""
	v.Password = ""
	v.Operator = ""
	v.EmployeeName = ""
	v.Email = ""
	v.Account = ""
	v.Language = ""
	v.SecMobileAreaCode = ""
	v.EmployeeNumber = ""
	v.SecMobile = ""
	v.Namespace = ""
	v.Avatar = ""
	v.TenantId = 0
	poolCreateTenantEmployeeAndAccountRequest.Put(v)
}

package mozi

import (
	"sync"
)

// AddOrUpdateTenantEmployeeAndAccountRequest 结构体
type AddOrUpdateTenantEmployeeAndAccountRequest struct {
	// 附加属性
	RequestMetaData string `json:"request_meta_data,omitempty" xml:"request_meta_data,omitempty"`
	// 昵称
	Nickname string `json:"nickname,omitempty" xml:"nickname,omitempty"`
	// 人员唯一编码，租户独一无二，比如工号
	EmployeeCode string `json:"employee_code,omitempty" xml:"employee_code,omitempty"`
	// 操作人
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 账号的所属公司，找BUC同学配置
	Namespace string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// 人员姓名
	EmployeeName string `json:"employee_name,omitempty" xml:"employee_name,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 账号名
	Account string `json:"account,omitempty" xml:"account,omitempty"`
	// 站点语言，中文ZH_CN,英文 EN
	Language string `json:"language,omitempty" xml:"language,omitempty"`
	// 手机国家码
	SecMobileAreaCode string `json:"sec_mobile_area_code,omitempty" xml:"sec_mobile_area_code,omitempty"`
	// 账号ID或者名字，能独一无二表示一个账号
	ReferId string `json:"refer_id,omitempty" xml:"refer_id,omitempty"`
	// 手机号码
	SecMobile string `json:"sec_mobile,omitempty" xml:"sec_mobile,omitempty"`
	// 身份证号码
	IdentityCardNum string `json:"identity_card_num,omitempty" xml:"identity_card_num,omitempty"`
	// 员工工号
	EmployeeNumber string `json:"employee_number,omitempty" xml:"employee_number,omitempty"`
	// 头像URL地址
	Avatar string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// 租户ID
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
}

var poolAddOrUpdateTenantEmployeeAndAccountRequest = sync.Pool{
	New: func() any {
		return new(AddOrUpdateTenantEmployeeAndAccountRequest)
	},
}

// GetAddOrUpdateTenantEmployeeAndAccountRequest() 从对象池中获取AddOrUpdateTenantEmployeeAndAccountRequest
func GetAddOrUpdateTenantEmployeeAndAccountRequest() *AddOrUpdateTenantEmployeeAndAccountRequest {
	return poolAddOrUpdateTenantEmployeeAndAccountRequest.Get().(*AddOrUpdateTenantEmployeeAndAccountRequest)
}

// ReleaseAddOrUpdateTenantEmployeeAndAccountRequest 释放AddOrUpdateTenantEmployeeAndAccountRequest
func ReleaseAddOrUpdateTenantEmployeeAndAccountRequest(v *AddOrUpdateTenantEmployeeAndAccountRequest) {
	v.RequestMetaData = ""
	v.Nickname = ""
	v.EmployeeCode = ""
	v.Operator = ""
	v.Namespace = ""
	v.EmployeeName = ""
	v.Email = ""
	v.Account = ""
	v.Language = ""
	v.SecMobileAreaCode = ""
	v.ReferId = ""
	v.SecMobile = ""
	v.IdentityCardNum = ""
	v.EmployeeNumber = ""
	v.Avatar = ""
	v.TenantId = 0
	poolAddOrUpdateTenantEmployeeAndAccountRequest.Put(v)
}

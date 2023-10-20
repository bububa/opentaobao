package mozi

import (
	"sync"
)

// TenantAdmin 结构体
type TenantAdmin struct {
	// 人员名
	EmployeeName string `json:"employee_name,omitempty" xml:"employee_name,omitempty"`
	// 人员code
	EmployeeCode string `json:"employee_code,omitempty" xml:"employee_code,omitempty"`
	// 是否为管理员
	Primary bool `json:"primary,omitempty" xml:"primary,omitempty"`
}

var poolTenantAdmin = sync.Pool{
	New: func() any {
		return new(TenantAdmin)
	},
}

// GetTenantAdmin() 从对象池中获取TenantAdmin
func GetTenantAdmin() *TenantAdmin {
	return poolTenantAdmin.Get().(*TenantAdmin)
}

// ReleaseTenantAdmin 释放TenantAdmin
func ReleaseTenantAdmin(v *TenantAdmin) {
	v.EmployeeName = ""
	v.EmployeeCode = ""
	v.Primary = false
	poolTenantAdmin.Put(v)
}

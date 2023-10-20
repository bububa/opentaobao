package mozi

import (
	"sync"
)

// EmployeeAccount 结构体
type EmployeeAccount struct {
	// 员工code
	EmployeeCode string `json:"employee_code,omitempty" xml:"employee_code,omitempty"`
	// 账号类型
	AccountNamespace string `json:"account_namespace,omitempty" xml:"account_namespace,omitempty"`
	// 账号名
	AccountCode string `json:"account_code,omitempty" xml:"account_code,omitempty"`
	// 账号ID
	AccountId int64 `json:"account_id,omitempty" xml:"account_id,omitempty"`
}

var poolEmployeeAccount = sync.Pool{
	New: func() any {
		return new(EmployeeAccount)
	},
}

// GetEmployeeAccount() 从对象池中获取EmployeeAccount
func GetEmployeeAccount() *EmployeeAccount {
	return poolEmployeeAccount.Get().(*EmployeeAccount)
}

// ReleaseEmployeeAccount 释放EmployeeAccount
func ReleaseEmployeeAccount(v *EmployeeAccount) {
	v.EmployeeCode = ""
	v.AccountNamespace = ""
	v.AccountCode = ""
	v.AccountId = 0
	poolEmployeeAccount.Put(v)
}

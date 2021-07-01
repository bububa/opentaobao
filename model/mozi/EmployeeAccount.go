package mozi

// EmployeeAccount 结构体
type EmployeeAccount struct {
	// 账号ID
	AccountId int64 `json:"account_id,omitempty" xml:"account_id,omitempty"`
	// 员工code
	EmployeeCode string `json:"employee_code,omitempty" xml:"employee_code,omitempty"`
	// 账号类型
	AccountNamespace string `json:"account_namespace,omitempty" xml:"account_namespace,omitempty"`
	// 账号名
	AccountCode string `json:"account_code,omitempty" xml:"account_code,omitempty"`
}

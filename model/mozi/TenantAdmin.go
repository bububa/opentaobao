package mozi

// TenantAdmin 结构体
type TenantAdmin struct {
	// 人员名
	EmployeeName string `json:"employee_name,omitempty" xml:"employee_name,omitempty"`
	// 人员code
	EmployeeCode string `json:"employee_code,omitempty" xml:"employee_code,omitempty"`
	// 是否为管理员
	Primary bool `json:"primary,omitempty" xml:"primary,omitempty"`
}

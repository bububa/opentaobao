package mozi

// GetEmployeeTenantAdminInfoRequest 结构体
type GetEmployeeTenantAdminInfoRequest struct {
	// 租户ID
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
	// 人员code
	EmployeeCode string `json:"employee_code,omitempty" xml:"employee_code,omitempty"`
}

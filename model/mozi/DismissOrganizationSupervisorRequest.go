package mozi

// DismissOrganizationSupervisorRequest 
type DismissOrganizationSupervisorRequest struct {
    // 员工 Code 列表
    EmployeeCodes   []string `json:"employee_codes,omitempty" xml:"employee_codes>string,omitempty"`
    // 组织 Code
    OrganizationCode   string `json:"organization_code,omitempty" xml:"organization_code,omitempty"`
    // 租户 ID
    TenantId   int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
    // 操作人
    Operator   string `json:"operator,omitempty" xml:"operator,omitempty"`
}

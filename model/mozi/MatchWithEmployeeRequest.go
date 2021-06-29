package mozi

// MatchWithEmployeeRequest 
type MatchWithEmployeeRequest struct {
    // 人员code
    EmployeeCodes   []string `json:"employee_codes,omitempty" xml:"employee_codes>string,omitempty"`
    // 租户id
    TenantId   int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
    // 组code
    GroupCode   string `json:"group_code,omitempty" xml:"group_code,omitempty"`
}

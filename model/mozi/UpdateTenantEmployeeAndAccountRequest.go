package mozi

// UpdateTenantEmployeeAndAccountRequest 
type UpdateTenantEmployeeAndAccountRequest struct {
    // 员工姓名
    EmployeeName   string `json:"employee_name,omitempty" xml:"employee_name,omitempty"`
    // 证件号码
    CertificateCode   string `json:"certificate_code,omitempty" xml:"certificate_code,omitempty"`
    // 员工基本属性
    EmployeeBaseProperties   *EmployeeBaseProperties `json:"employee_base_properties,omitempty" xml:"employee_base_properties,omitempty"`
    // 操作人
    Operator   string `json:"operator,omitempty" xml:"operator,omitempty"`
    // 请求附加消息
    RequestMetaData   string `json:"request_meta_data,omitempty" xml:"request_meta_data,omitempty"`
    // 员工CODE
    EmployeeCode   string `json:"employee_code,omitempty" xml:"employee_code,omitempty"`
    // 账号ID
    AccountId   int64 `json:"account_id,omitempty" xml:"account_id,omitempty"`
    // 租户ID
    TenantId   int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
    // 证件类型
    CertificateType   string `json:"certificate_type,omitempty" xml:"certificate_type,omitempty"`
}

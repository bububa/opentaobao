package mozi

// RemoveTenantEmployeeAndAccountRequest 
type RemoveTenantEmployeeAndAccountRequest struct {

    // 租户ID
    
    TenantId   int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
    

    // 附加属性
    
    RequestMetaData   string `json:"request_meta_data,omitempty" xml:"request_meta_data,omitempty"`
    

    // 人员CODE
    
    EmployeeCode   string `json:"employee_code,omitempty" xml:"employee_code,omitempty"`
    

    // 操作人
    
    Operator   string `json:"operator,omitempty" xml:"operator,omitempty"`
    

    // 账号ID
    
    AccountId   int64 `json:"account_id,omitempty" xml:"account_id,omitempty"`
    

}

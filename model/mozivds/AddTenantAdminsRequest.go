package mozivds

// AddTenantAdminsRequest 
type AddTenantAdminsRequest struct {

    // 人员code
    
    EmployeeCodes   []string `json:"employee_codes,omitempty" xml:"employee_codes>string,omitempty"`
    

    // 是否主管理员
    
    PrimaryAdmin   bool `json:"primary_admin,omitempty" xml:"primary_admin,omitempty"`
    

    // 租户Id
    
    TenantId   int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
    

    // 操作人
    
    Operator   string `json:"operator,omitempty" xml:"operator,omitempty"`
    

    // 请求元数据
    
    RequestMetaData   string `json:"request_meta_data,omitempty" xml:"request_meta_data,omitempty"`
    

}

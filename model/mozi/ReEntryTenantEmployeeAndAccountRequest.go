package mozi

// ReEntryTenantEmployeeAndAccountRequest 
type ReEntryTenantEmployeeAndAccountRequest struct {
    // 租户ID
    TenantId   int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
    // 证件号码
    CertificateCode   string `json:"certificate_code,omitempty" xml:"certificate_code,omitempty"`
    // 请求附加消息
    RequestMetaData   string `json:"request_meta_data,omitempty" xml:"request_meta_data,omitempty"`
    // 昵称
    Nickname   string `json:"nickname,omitempty" xml:"nickname,omitempty"`
    // 证件类型
    CertificateType   string `json:"certificate_type,omitempty" xml:"certificate_type,omitempty"`
    // 员工CODE
    EmployeeCode   string `json:"employee_code,omitempty" xml:"employee_code,omitempty"`
    // 密码
    Password   string `json:"password,omitempty" xml:"password,omitempty"`
    // 操作人
    Operator   string `json:"operator,omitempty" xml:"operator,omitempty"`
    // 员工姓名
    EmployeeName   string `json:"employee_name,omitempty" xml:"employee_name,omitempty"`
    // 账号平台
    Namespace   string `json:"namespace,omitempty" xml:"namespace,omitempty"`
    // 邮箱
    Email   string `json:"email,omitempty" xml:"email,omitempty"`
    // 账号名
    Account   string `json:"account,omitempty" xml:"account,omitempty"`
    // 语言
    Language   string `json:"language,omitempty" xml:"language,omitempty"`
    // 手机区号
    SecMobileAreaCode   string `json:"sec_mobile_area_code,omitempty" xml:"sec_mobile_area_code,omitempty"`
    // 工号
    EmployeeNumber   string `json:"employee_number,omitempty" xml:"employee_number,omitempty"`
    // 信任手机
    SecMobile   string `json:"sec_mobile,omitempty" xml:"sec_mobile,omitempty"`
    // 头像URL
    Avatar   string `json:"avatar,omitempty" xml:"avatar,omitempty"`
}

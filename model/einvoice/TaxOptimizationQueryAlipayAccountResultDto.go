package einvoice

// TaxOptimizationQueryAlipayAccountResultDto 
type TaxOptimizationQueryAlipayAccountResultDto struct {
    // 账号类型
    AccountTypeEnum   int64 `json:"account_type_enum,omitempty" xml:"account_type_enum,omitempty"`
    // 是否正常发薪
    CanPay   bool `json:"can_pay,omitempty" xml:"can_pay,omitempty"`
    // 认证类型
    CertTypeEnum   string `json:"cert_type_enum,omitempty" xml:"cert_type_enum,omitempty"`
    // 是否实人认证
    Certified   bool `json:"certified,omitempty" xml:"certified,omitempty"`
    // 账号状态
    EnableStatusEnum   string `json:"enable_status_enum,omitempty" xml:"enable_status_enum,omitempty"`
    // 发薪账号
    PaySalaryAlipayAccount   string `json:"pay_salary_alipay_account,omitempty" xml:"pay_salary_alipay_account,omitempty"`
}

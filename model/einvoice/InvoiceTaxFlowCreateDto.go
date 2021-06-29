package einvoice

// InvoiceTaxFlowCreateDto 
type InvoiceTaxFlowCreateDto struct {
    // 企业税务信息
    InvoiceCompany   *InvoiceCompanyDto `json:"invoice_company,omitempty" xml:"invoice_company,omitempty"`
    // 联系人信息
    InvoiceContact   *InvoiceContactDto `json:"invoice_contact,omitempty" xml:"invoice_contact,omitempty"`
    // 外部业务方创建税控开通工单的唯一幂等ID（即：相同outer_id 会被视为同一个请求，被幂等处理）, 由业务方自己生成。  只能由字母数字组成
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    // 业务平台code, 由阿里发票小二分配
    PlatformCode   string `json:"platform_code,omitempty" xml:"platform_code,omitempty"`
    // 业务平台商户ID
    PlatformUserId   string `json:"platform_user_id,omitempty" xml:"platform_user_id,omitempty"`
    // 税控产品，产品码由中台定义。
    ProductCode   string `json:"product_code,omitempty" xml:"product_code,omitempty"`
    // 服务的有效天数，单位为天。  阿里发票服务周期计算规则为：服务起始时间=部署完成时的系统时间，服务截止时间=服务起始时间+serviceValidDays
    ServiceValidDays   int64 `json:"service_valid_days,omitempty" xml:"service_valid_days,omitempty"`
}

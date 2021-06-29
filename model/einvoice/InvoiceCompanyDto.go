package einvoice

// InvoiceCompanyDto 
type InvoiceCompanyDto struct {

    // 企业所在区
    
    Area   string `json:"area,omitempty" xml:"area,omitempty"`
    

    // 开户行账号，注意：开户行账号与名称必须拆开2个字段
    
    BankAccountId   string `json:"bank_account_id,omitempty" xml:"bank_account_id,omitempty"`
    

    // 开户行名称，开户行账号加名称不超出100字符注意：开户行账号与名称必须拆开2个字段
    
    BankName   string `json:"bank_name,omitempty" xml:"bank_name,omitempty"`
    

    // 企业所在城市。  请提交完整的城市名称，正确示例：杭州市，错误示例：杭州
    
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    

    // 企业名称
    
    CompanyName   string `json:"company_name,omitempty" xml:"company_name,omitempty"`
    

    // 企业类型，可选值：  一般纳税人：1;  小规模纳税人：2;  起征点以下纳税人：3;
    
    CompanyType   int64 `json:"company_type,omitempty" xml:"company_type,omitempty"`
    

    // 默认商品名称
    
    DefaultItemName   string `json:"default_item_name,omitempty" xml:"default_item_name,omitempty"`
    

    // 默认税收分类编码
    
    DefaultTaxCode   string `json:"default_tax_code,omitempty" xml:"default_tax_code,omitempty"`
    

    // 税率，格式为小数
    
    DefaultTaxRate   string `json:"default_tax_rate,omitempty" xml:"default_tax_rate,omitempty"`
    

    // 企业地址
    
    DetailedAddress   string `json:"detailed_address,omitempty" xml:"detailed_address,omitempty"`
    

    // 企业电话
    
    InvoicePhone   string `json:"invoice_phone,omitempty" xml:"invoice_phone,omitempty"`
    

    // 复核人
    
    PayeeChecker   string `json:"payee_checker,omitempty" xml:"payee_checker,omitempty"`
    

    // 默认开票人
    
    PayeeOperator   string `json:"payee_operator,omitempty" xml:"payee_operator,omitempty"`
    

    // 收款人
    
    PayeeReceiver   string `json:"payee_receiver,omitempty" xml:"payee_receiver,omitempty"`
    

    // 销方纳税人识别号
    
    PayeeRegisterNo   string `json:"payee_register_no,omitempty" xml:"payee_register_no,omitempty"`
    

    // 企业所在省/直辖市
    
    Province   string `json:"province,omitempty" xml:"province,omitempty"`
    

    // 0税率标识，只有税率为0的情况才有值，0=出口零税率，1=免税，2=不征收，3=普通零税率
    
    ZeroTaxRateFlag   string `json:"zero_tax_rate_flag,omitempty" xml:"zero_tax_rate_flag,omitempty"`
    

}

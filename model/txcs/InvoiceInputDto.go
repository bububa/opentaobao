package txcs

// InvoiceInputDto 
type InvoiceInputDto struct {

    // 结算公司
    
    SettlementCompanyCode   string `json:"settlement_company_code,omitempty" xml:"settlement_company_code,omitempty"`
    

    // 发票日期
    
    InvoiceDate   string `json:"invoice_date,omitempty" xml:"invoice_date,omitempty"`
    

    // 发品编号
    
    InvoiceCode   string `json:"invoice_code,omitempty" xml:"invoice_code,omitempty"`
    

    // 操作人员
    
    OperatorName   string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
    

    // 税率
    
    TaxRate   string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
    

    // 总金额
    
    TotalAmount   string `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
    

    // 幂等ID
    
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
    

    // 发票类型
    
    InvoiceType   string `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
    

    // 税率
    
    UntaxAmount   string `json:"untax_amount,omitempty" xml:"untax_amount,omitempty"`
    

    // 币种
    
    Currency   string `json:"currency,omitempty" xml:"currency,omitempty"`
    

    // 发票编码
    
    InvoiceNo   string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
    

    // 税票金额
    
    TaxAmount   string `json:"tax_amount,omitempty" xml:"tax_amount,omitempty"`
    

    // 操作人员ID
    
    OperatorId   string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
    

}

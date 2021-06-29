package txcs

// StatementBillDto 
type StatementBillDto struct {

    // 结算主体编码
    
    SettlementCompanyCode   string `json:"settlement_company_code,omitempty" xml:"settlement_company_code,omitempty"`
    

    // 经营属性描述
    
    OperateTypeDesc   string `json:"operate_type_desc,omitempty" xml:"operate_type_desc,omitempty"`
    

    // 经营属性
    
    OperateType   string `json:"operate_type,omitempty" xml:"operate_type,omitempty"`
    

    // 供应商名称
    
    BusinessPartnerName   string `json:"business_partner_name,omitempty" xml:"business_partner_name,omitempty"`
    

    // 供应商编码
    
    BusinessPartnerCode   string `json:"business_partner_code,omitempty" xml:"business_partner_code,omitempty"`
    

    // 商家编码
    
    MerchantCode   string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
    

    // 对账单号
    
    StatementBillCode   string `json:"statement_bill_code,omitempty" xml:"statement_bill_code,omitempty"`
    

    // 结算主体名称
    
    SettlementCompanyName   string `json:"settlement_company_name,omitempty" xml:"settlement_company_name,omitempty"`
    

    // 账单日
    
    BillDay   string `json:"bill_day,omitempty" xml:"bill_day,omitempty"`
    

    // 订单状态
    
    BillStatus   string `json:"bill_status,omitempty" xml:"bill_status,omitempty"`
    

    // 订单状态描述
    
    BillStatusDesc   string `json:"bill_status_desc,omitempty" xml:"bill_status_desc,omitempty"`
    

    // 应结算总金额
    
    SettlementTotalAmount   string `json:"settlement_total_amount,omitempty" xml:"settlement_total_amount,omitempty"`
    

    // 应开票总金额
    
    InvoiceTotalAmount   string `json:"invoice_total_amount,omitempty" xml:"invoice_total_amount,omitempty"`
    

    // 结算周期描述
    
    SettlePeridDesc   string `json:"settle_perid_desc,omitempty" xml:"settle_perid_desc,omitempty"`
    

    // 币种
    
    Currency   string `json:"currency,omitempty" xml:"currency,omitempty"`
    

}

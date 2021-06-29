package einvoice

// ResultList 
type ResultList struct {

    // 付款方税号
    
    PayeeRegisterNo   string `json:"payee_register_no,omitempty" xml:"payee_register_no,omitempty"`
    

    // 付款方平台
    
    Platform   string `json:"platform,omitempty" xml:"platform,omitempty"`
    

    // 订单id
    
    OrderId   string `json:"order_id,omitempty" xml:"order_id,omitempty"`
    

    // 开票金额
    
    SumPrice   string `json:"sum_price,omitempty" xml:"sum_price,omitempty"`
    

    // 开票明细列表
    
    InvoiceItems   []InvoiceItems `json:"invoice_items,omitempty" xml:"invoice_items,omitempty"`
    

    // seriNo
    
    SeriNo   string `json:"seri_no,omitempty" xml:"seri_no,omitempty"`
    

    // invoiceStatus
    
    InvoiceStatus   int64 `json:"invoice_status,omitempty" xml:"invoice_status,omitempty"`
    

}

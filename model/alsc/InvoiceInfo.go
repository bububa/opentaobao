package alsc

// InvoiceInfo 
/* model for simplify = false
type InvoiceInfo struct {

    // 发票抬头
    
    Invoice   string `json:"invoice,omitempty"`
    

    // 发票类型：  个人-PERSONAL公司-COMPANY
    
    InvoiceType   string `json:"invoice_type,omitempty"`
    

    // 纳税人识别码
    
    TaxPayerId   string `json:"tax_payer_id,omitempty"`
    

}
*/

// InvoiceInfo 
type InvoiceInfo struct {

    // 发票抬头
    Invoice   string `json:"invoice,omitempty"`

    // 发票类型：  个人-PERSONAL公司-COMPANY
    InvoiceType   string `json:"invoice_type,omitempty"`

    // 纳税人识别码
    TaxPayerId   string `json:"tax_payer_id,omitempty"`

}

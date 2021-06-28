package wms

// Invoinceconfirm 
/* model for simplify = false
type Invoinceconfirm struct {

    // Erp发票ID
    
    BillId   string `json:"bill_id,omitempty"`
    

    // 发票号
    
    InvoiceNumber   string `json:"invoice_number,omitempty"`
    

    // 发票代码
    
    InvoiceCode   string `json:"invoice_code,omitempty"`
    

}
*/

// Invoinceconfirm 
type Invoinceconfirm struct {

    // Erp发票ID
    BillId   string `json:"bill_id,omitempty"`

    // 发票号
    InvoiceNumber   string `json:"invoice_number,omitempty"`

    // 发票代码
    InvoiceCode   string `json:"invoice_code,omitempty"`

}

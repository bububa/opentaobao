package flight

// Tax 
/* model for simplify = false
type Tax struct {

    // 税费编码
    
    TaxCode   string `json:"tax_code,omitempty"`
    

    // 税费金额
    
    Amount   int64 `json:"amount,omitempty"`
    

    // 不可退税费
    
    NotRefundAmount   int64 `json:"not_refund_amount,omitempty"`
    

}
*/

// Tax 
type Tax struct {

    // 税费编码
    TaxCode   string `json:"tax_code,omitempty"`

    // 税费金额
    Amount   int64 `json:"amount,omitempty"`

    // 不可退税费
    NotRefundAmount   int64 `json:"not_refund_amount,omitempty"`

}

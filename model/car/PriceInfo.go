package car

// PriceInfo 
/* model for simplify = false
type PriceInfo struct {

    // 总费用，折后金额总费用，折后金额  eg:203.00
    
    TotalPrice   string `json:"total_price,omitempty"`
    

    // 原价，如果订单有折扣这里为折扣前的价格，如果没有折扣和totalPrice字段保持一致
    
    OriginalPrice   string `json:"original_price,omitempty"`
    

    // 费用明细
    
    Detail  struct {
        DetailPriceInfo  []DetailPriceInfo `json:"detail_price_info,omitempty"`
    } `json:"detail,omitempty"`
    

}
*/

// PriceInfo 
type PriceInfo struct {

    // 总费用，折后金额总费用，折后金额  eg:203.00
    TotalPrice   string `json:"total_price,omitempty"`

    // 原价，如果订单有折扣这里为折扣前的价格，如果没有折扣和totalPrice字段保持一致
    OriginalPrice   string `json:"original_price,omitempty"`

    // 费用明细
    Detail   []DetailPriceInfo `json:"detail,omitempty"`

}

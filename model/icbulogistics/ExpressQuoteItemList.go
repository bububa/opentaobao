package icbulogistics

// ExpressQuoteItemList 
type ExpressQuoteItemList struct {

    // 费用编码
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

    // 数量
    
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    

    // 价格信息
    
    SalesAmount   *Money `json:"sales_amount,omitempty" xml:"sales_amount,omitempty"`
    

    // 费用名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 费用描述
    
    ChargeDesc   string `json:"charge_desc,omitempty" xml:"charge_desc,omitempty"`
    

    // 币种
    
    Currency   string `json:"currency,omitempty" xml:"currency,omitempty"`
    

    // 费用类型
    
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    

}

package trade

// ExpandCardInfo 
type ExpandCardInfo struct {

    // 买卡订单本金
    
    BasicPrice   string `json:"basic_price,omitempty" xml:"basic_price,omitempty"`
    

    // 买卡订单权益金
    
    ExpandPrice   string `json:"expand_price,omitempty" xml:"expand_price,omitempty"`
    

    // 用卡订单使用的本金
    
    BasicPriceUsed   string `json:"basic_price_used,omitempty" xml:"basic_price_used,omitempty"`
    

    // 用卡订单使用的权益金
    
    ExpandPriceUsed   string `json:"expand_price_used,omitempty" xml:"expand_price_used,omitempty"`
    

}

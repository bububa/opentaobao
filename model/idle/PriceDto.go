package idle

// PriceDto 
type PriceDto struct {

    // 押金，单位分
    
    Deposit   int64 `json:"deposit,omitempty" xml:"deposit,omitempty"`
    

    // 市场价，单位分
    
    OriginalPrice   int64 `json:"original_price,omitempty" xml:"original_price,omitempty"`
    

    // 买断价，单位分
    
    ReservePrice   int64 `json:"reserve_price,omitempty" xml:"reserve_price,omitempty"`
    

    // 日结算价格，单位分
    
    SettlePricePerDay   int64 `json:"settle_price_per_day,omitempty" xml:"settle_price_per_day,omitempty"`
    

    // 星级
    
    Star   string `json:"star,omitempty" xml:"star,omitempty"`
    

}

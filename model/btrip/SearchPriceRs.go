package btrip

// SearchPriceRs 
type SearchPriceRs struct {

    // 原始销售价，活动前的价格（单位：元）
    
    OriginalSellPrice   int64 `json:"original_sell_price,omitempty" xml:"original_sell_price,omitempty"`
    

    // 销售价(单位：元)
    
    SellPrice   int64 `json:"sell_price,omitempty" xml:"sell_price,omitempty"`
    

    // 参考税(单位：元)
    
    Tax   int64 `json:"tax,omitempty" xml:"tax,omitempty"`
    

}

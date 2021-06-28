package travel

// Prices 
type Prices struct {

    // 日期
    
    Date   string `json:"date,omitempty" xml:"date,omitempty"`
    

    // 外部商家团期ID
    
    OuterPriceId   string `json:"outer_price_id,omitempty" xml:"outer_price_id,omitempty"`
    

    // 价格
    
    Price   int64 `json:"price,omitempty" xml:"price,omitempty"`
    

    // 库存
    
    Stock   int64 `json:"stock,omitempty" xml:"stock,omitempty"`
    

    // 价格类型。price_type 取：1-成人价，2-儿童价，3-单房差
    
    PriceType   int64 `json:"price_type,omitempty" xml:"price_type,omitempty"`
    

}

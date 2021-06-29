package btrip

// BtripHotelPromotionDetailDto 
type BtripHotelPromotionDetailDto struct {

    // 优惠项名称
    
    PromotionName   string `json:"promotion_name,omitempty" xml:"promotion_name,omitempty"`
    

    // 优惠金额
    
    PromotionPrice   int64 `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
    

    // 优惠类型
    
    PromotionType   int64 `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
    

}

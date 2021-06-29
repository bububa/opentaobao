package xhotelonlineorder

// OrderSettleDo 
type OrderSettleDo struct {

    // 平台优惠金额
    
    PlatPromotionPrice   int64 `json:"plat_promotion_price,omitempty" xml:"plat_promotion_price,omitempty"`
    

    // 其他费用
    
    OtherFee   int64 `json:"other_fee,omitempty" xml:"other_fee,omitempty"`
    

    // 房费
    
    RoomPrice   int64 `json:"room_price,omitempty" xml:"room_price,omitempty"`
    

    // 结账金额
    
    OrderSettlePrice   int64 `json:"order_settle_price,omitempty" xml:"order_settle_price,omitempty"`
    

    // 实际结账日历
    
    SettleDailyPrice   *HbsDailyPriceDo `json:"settle_daily_price,omitempty" xml:"settle_daily_price,omitempty"`
    

    // 卖家优惠金额
    
    PromotionPrice   int64 `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
    

}

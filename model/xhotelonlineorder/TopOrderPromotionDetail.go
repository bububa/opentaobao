package xhotelonlineorder

// TopOrderPromotionDetail 
type TopOrderPromotionDetail struct {

    // 返现金额
    
    CashBackAmount   int64 `json:"cash_back_amount,omitempty" xml:"cash_back_amount,omitempty"`
    

    // 卖家立减
    
    SellerDecrease   int64 `json:"seller_decrease,omitempty" xml:"seller_decrease,omitempty"`
    

    // topHotelPromotions
    
    TopHotelPromotions   []TopHotelPromotion `json:"top_hotel_promotions,omitempty" xml:"top_hotel_promotions,omitempty"`
    

    // 立减金额
    
    Decrease   int64 `json:"decrease,omitempty" xml:"decrease,omitempty"`
    

    // 平台立减金额
    
    PlatformDecrease   int64 `json:"platform_decrease,omitempty" xml:"platform_decrease,omitempty"`
    

}

package promotion

// BuyerCouponInfo 
type BuyerCouponInfo struct {

    // 券名称
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 卖家ID
    
    SellerId   string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
    

    // 状态名称
    
    StatusName   string `json:"status_name,omitempty" xml:"status_name,omitempty"`
    

    // 券类型
    
    CouponTypeName   string `json:"coupon_type_name,omitempty" xml:"coupon_type_name,omitempty"`
    

    // 面额
    
    Amount   int64 `json:"amount,omitempty" xml:"amount,omitempty"`
    

    // 券生效时间
    
    StartTime   string `json:"start_time,omitempty" xml:"start_time,omitempty"`
    

    // 券失效时间
    
    EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
    

    // 满足金额阀值  如订单满多少元才可用
    
    StartFee   int64 `json:"start_fee,omitempty" xml:"start_fee,omitempty"`
    

    // mtop 店铺链接
    
    Url   string `json:"url,omitempty" xml:"url,omitempty"`
    

    // 卖家名称
    
    SellerNick   string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
    

    // 券对外ID
    
    SpreadId   string `json:"spread_id,omitempty" xml:"spread_id,omitempty"`
    

    // 券类型
    
    CouponType   int64 `json:"coupon_type,omitempty" xml:"coupon_type,omitempty"`
    

    // 店铺名称
    
    ShopName   string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
    

    // 状态信息
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

}

package promotion

// SellerCouponDetail 
/* model for simplify = false
type SellerCouponDetail struct {

    // 券名称
    
    Title   string `json:"title,omitempty"`
    

    // 卖家ID
    
    SellerId   string `json:"seller_id,omitempty"`
    

    // 状态名称
    
    StatusName   string `json:"status_name,omitempty"`
    

    // 券类型
    
    CouponTypeName   string `json:"coupon_type_name,omitempty"`
    

    // 面额
    
    Amount   int64 `json:"amount,omitempty"`
    

    // 设置发券数量总数
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

    // 券生效时间
    
    StartTime   string `json:"start_time,omitempty"`
    

    // 券失效时间
    
    EndTime   string `json:"end_time,omitempty"`
    

    // 满足金额阀值  如订单满多少元才可用
    
    StartFee   int64 `json:"start_fee,omitempty"`
    

    // mtop 店铺链接
    
    Url   string `json:"url,omitempty"`
    

    // 卖家名称
    
    SellerNick   string `json:"seller_nick,omitempty"`
    

    // 商品优惠券会有商品id集合
    
    ItemIds  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"item_ids,omitempty"`
    

    // 券对外ID
    
    SpreadId   string `json:"spread_id,omitempty"`
    

    // 券类型
    
    CouponType   int64 `json:"coupon_type,omitempty"`
    

    // 店铺名称
    
    ShopName   string `json:"shop_name,omitempty"`
    

    // 状态信息
    
    Status   int64 `json:"status,omitempty"`
    

}
*/

// SellerCouponDetail 
type SellerCouponDetail struct {

    // 券名称
    Title   string `json:"title,omitempty"`

    // 卖家ID
    SellerId   string `json:"seller_id,omitempty"`

    // 状态名称
    StatusName   string `json:"status_name,omitempty"`

    // 券类型
    CouponTypeName   string `json:"coupon_type_name,omitempty"`

    // 面额
    Amount   int64 `json:"amount,omitempty"`

    // 设置发券数量总数
    TotalCount   int64 `json:"total_count,omitempty"`

    // 券生效时间
    StartTime   string `json:"start_time,omitempty"`

    // 券失效时间
    EndTime   string `json:"end_time,omitempty"`

    // 满足金额阀值  如订单满多少元才可用
    StartFee   int64 `json:"start_fee,omitempty"`

    // mtop 店铺链接
    Url   string `json:"url,omitempty"`

    // 卖家名称
    SellerNick   string `json:"seller_nick,omitempty"`

    // 商品优惠券会有商品id集合
    ItemIds   []int64 `json:"item_ids,omitempty"`

    // 券对外ID
    SpreadId   string `json:"spread_id,omitempty"`

    // 券类型
    CouponType   int64 `json:"coupon_type,omitempty"`

    // 店铺名称
    ShopName   string `json:"shop_name,omitempty"`

    // 状态信息
    Status   int64 `json:"status,omitempty"`

}

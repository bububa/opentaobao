package alsc

// ChargePreCheckOpenReq 
/* model for simplify = false
type ChargePreCheckOpenReq struct {

    // 品牌ID(和outbrandid不能同时为空)
    
    BrandId   string `json:"brand_id,omitempty"`
    

    // 卡Id，用于会员卡、礼品卡校验
    
    CardId   string `json:"card_id,omitempty"`
    

    // 会员ID，用于会员卡校验，会员卡必填
    
    CustomerId   string `json:"customer_id,omitempty"`
    

    // 门店ID(和outshopid不能同时为空)
    
    ShopId   string `json:"shop_id,omitempty"`
    

    // 外部门店ID(不能和shopid同时为空)
    
    OutShopId   string `json:"out_shop_id,omitempty"`
    

    // 外部品牌ID(不能和brandid同时为空)
    
    OutBrandId   string `json:"out_brand_id,omitempty"`
    

}
*/

// ChargePreCheckOpenReq 
type ChargePreCheckOpenReq struct {

    // 品牌ID(和outbrandid不能同时为空)
    BrandId   string `json:"brand_id,omitempty"`

    // 卡Id，用于会员卡、礼品卡校验
    CardId   string `json:"card_id,omitempty"`

    // 会员ID，用于会员卡校验，会员卡必填
    CustomerId   string `json:"customer_id,omitempty"`

    // 门店ID(和outshopid不能同时为空)
    ShopId   string `json:"shop_id,omitempty"`

    // 外部门店ID(不能和shopid同时为空)
    OutShopId   string `json:"out_shop_id,omitempty"`

    // 外部品牌ID(不能和brandid同时为空)
    OutBrandId   string `json:"out_brand_id,omitempty"`

}

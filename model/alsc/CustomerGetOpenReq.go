package alsc

// CustomerGetOpenReq 
type CustomerGetOpenReq struct {

    // saas品牌id
    BrandId   string `json:"brand_id,omitempty"`

    // 手机号
    Mobile   string `json:"mobile,omitempty"`

    // 外部id
    OuterId   string `json:"outer_id,omitempty"`

    // 外部id类型，wechat：微信openId  alipay：支付宝
    OuterType   string `json:"outer_type,omitempty"`

    // 物理卡id
    PhysicalCardId   string `json:"physical_card_id,omitempty"`

    // saas门店id
    ShopId   string `json:"shop_id,omitempty"`

    // 顾客id
    CustomerId   string `json:"customer_id,omitempty"`

}

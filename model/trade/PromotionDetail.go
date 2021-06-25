package trade

// PromotionDetail 
type PromotionDetail struct {

    // 优惠id，(由营销工具id、优惠活动id和优惠详情id组成，结构为：营销工具id-优惠活动id_优惠详情id，如mjs-123024_211143）
    PromotionId   string `json:"promotion_id,omitempty"`

    // 优惠活动的描述
    PromotionDesc   string `json:"promotion_desc,omitempty"`

    // 满就送礼物的礼物数量
    GiftItemNum   string `json:"gift_item_num,omitempty"`

    // 赠品的宝贝id
    GiftItemId   string `json:"gift_item_id,omitempty"`

    // 满就送商品时，所送商品的名称
    GiftItemName   string `json:"gift_item_name,omitempty"`

    // 优惠金额（免运费、限时打折时为空）,单位：元
    DiscountFee   string `json:"discount_fee,omitempty"`

    // 优惠信息的名称
    PromotionName   string `json:"promotion_name,omitempty"`

    // 交易的主订单或子订单号
    Id   string `json:"id,omitempty"`

}

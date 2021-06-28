package iot

// PromotionInItem 
/* model for simplify = false
type PromotionInItem struct {

    // idValue的值
    
    PromotionId   string `json:"promotion_id,omitempty"`
    

    // 优惠展示名称
    
    Name   string `json:"name,omitempty"`
    

    // 优惠折后价格
    
    ItemPromoPrice   string `json:"item_promo_price,omitempty"`
    

    // sku价格列表
    
    SkuPriceList   string `json:"sku_price_list,omitempty"`
    

    // 优惠描述
    
    Desc   string `json:"desc,omitempty"`
    

    // 优惠开始时间
    
    StartTime   string `json:"start_time,omitempty"`
    

    // 优惠结束时间
    
    EndTime   string `json:"end_time,omitempty"`
    

    // 需要支付附加物，显示为+xxx。如：+20淘金币
    
    OtherNeed   string `json:"other_need,omitempty"`
    

    // 赠送东西。如：送10商城积分
    
    OtherSend   string `json:"other_send,omitempty"`
    

    // sku价格对应的id（保证二者顺序相同）
    
    SkuIdList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"sku_id_list,omitempty"`
    

}
*/

// PromotionInItem 
type PromotionInItem struct {

    // idValue的值
    PromotionId   string `json:"promotion_id,omitempty"`

    // 优惠展示名称
    Name   string `json:"name,omitempty"`

    // 优惠折后价格
    ItemPromoPrice   string `json:"item_promo_price,omitempty"`

    // sku价格列表
    SkuPriceList   string `json:"sku_price_list,omitempty"`

    // 优惠描述
    Desc   string `json:"desc,omitempty"`

    // 优惠开始时间
    StartTime   string `json:"start_time,omitempty"`

    // 优惠结束时间
    EndTime   string `json:"end_time,omitempty"`

    // 需要支付附加物，显示为+xxx。如：+20淘金币
    OtherNeed   string `json:"other_need,omitempty"`

    // 赠送东西。如：送10商城积分
    OtherSend   string `json:"other_send,omitempty"`

    // sku价格对应的id（保证二者顺序相同）
    SkuIdList   []string `json:"sku_id_list,omitempty"`

}

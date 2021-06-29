package tbk

// TaobaoTbkDgOptimusPromotionMapData 
type TaobaoTbkDgOptimusPromotionMapData struct {
    // 权益类型。1 有价券（需要购买的店铺券或单品券） 2 公开券（直接领取的店铺券或单品券） 3 妈妈券（妈妈渠道领取的店铺券或单品券） 4.品类券 （跨店可用券，可与1，2，3叠加）
    PromotionType   string `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
    // 优惠门槛类型： 1 满元 2 满件
    ConditionType   string `json:"condition_type,omitempty" xml:"condition_type,omitempty"`
    // 优惠类型： 1 减钱 2 打折
    DiscountType   string `json:"discount_type,omitempty" xml:"discount_type,omitempty"`
    // 权益信息-总量（权益初始库存量）
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    // 权益信息-剩余库存（权益剩余库存量）
    RemainCount   int64 `json:"remain_count,omitempty" xml:"remain_count,omitempty"`
    // 权益信息展示开始时间，精确到毫秒时间戳
    DisplayStartTime   string `json:"display_start_time,omitempty" xml:"display_start_time,omitempty"`
    // 权益信息展示结束时间，精确到毫秒时间戳
    DisplayEndTime   string `json:"display_end_time,omitempty" xml:"display_end_time,omitempty"`
    // 权益信息
    PromotionList   []PromotionList `json:"promotion_list,omitempty" xml:"promotion_list>promotion_list,omitempty"`
    // 权益扩展信息
    PromotionExtend   *PromotionExtend `json:"promotion_extend,omitempty" xml:"promotion_extend,omitempty"`
    // 店铺信息-卖家ID
    SellerId   string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
    // 店铺信息-卖家昵称
    Nick   string `json:"nick,omitempty" xml:"nick,omitempty"`
    // 店铺信息-店铺名称
    ShopTitle   string `json:"shop_title,omitempty" xml:"shop_title,omitempty"`
    // 店铺信息-店铺logo
    ShopPictureUrl   string `json:"shop_picture_url,omitempty" xml:"shop_picture_url,omitempty"`
}

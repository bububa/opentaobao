package wdk

// PromotionPriceDo 结构体
type PromotionPriceDo struct {
	// 促销说明
	PromotionReason string `json:"promotion_reason,omitempty" xml:"promotion_reason,omitempty"`
	// 店铺编码
	ShopCode string `json:"shop_code,omitempty" xml:"shop_code,omitempty"`
	// 促销类型
	PromotionType string `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
	// 促销开始时间
	PromotionStart string `json:"promotion_start,omitempty" xml:"promotion_start,omitempty"`
	// 促销赠品信息
	PromotionGiftInfo string `json:"promotion_gift_info,omitempty" xml:"promotion_gift_info,omitempty"`
	// 促销结束时间
	PromotionEnd string `json:"promotion_end,omitempty" xml:"promotion_end,omitempty"`
	// 商家编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 会员促销价结束时间
	MemberPromotionEndTime string `json:"member_promotion_end_time,omitempty" xml:"member_promotion_end_time,omitempty"`
	// 会员促销价开始时间
	MemberPromotionStartTime string `json:"member_promotion_start_time,omitempty" xml:"member_promotion_start_time,omitempty"`
	// 会员促销活动类型
	MemberPromotionType string `json:"member_promotion_type,omitempty" xml:"member_promotion_type,omitempty"`
	// 记录标识
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 店铺标识
	ShopId int64 `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 促销价格
	PromotionPrice int64 `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
	// 是否参加了促销
	IfPromotion int64 `json:"if_promotion,omitempty" xml:"if_promotion,omitempty"`
	// 会员促销价
	MemberPromotionPrice int64 `json:"member_promotion_price,omitempty" xml:"member_promotion_price,omitempty"`
}

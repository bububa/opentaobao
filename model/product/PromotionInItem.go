package product

// PromotionInItem 结构体
type PromotionInItem struct {
	// idValue的值
	PromotionId string `json:"promotion_id,omitempty" xml:"promotion_id,omitempty"`
	// 优惠展示名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 优惠折后价格
	ItemPromoPrice string `json:"item_promo_price,omitempty" xml:"item_promo_price,omitempty"`
	// sku价格列表
	SkuPriceList []string `json:"sku_price_list,omitempty" xml:"sku_price_list>string,omitempty"`
	// 优惠描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 优惠开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 优惠结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 需要支付附加物，显示为+xxx。如：+20淘金币
	OtherNeed string `json:"other_need,omitempty" xml:"other_need,omitempty"`
	// 赠送东西。如：送10商城积分
	OtherSend string `json:"other_send,omitempty" xml:"other_send,omitempty"`
	// sku价格对应的id（保证二者顺序相同）
	SkuIdList []string `json:"sku_id_list,omitempty" xml:"sku_id_list>string,omitempty"`
}

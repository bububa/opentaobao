package wdk

// PromotionSku 结构体
type PromotionSku struct {
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 采购进价
	PurchasePrice string `json:"purchase_price,omitempty" xml:"purchase_price,omitempty"`
	// 采购进价开始时间
	PurchasePriceStartTime string `json:"purchase_price_start_time,omitempty" xml:"purchase_price_start_time,omitempty"`
	// 采购进价结束时间
	PurchasePriceEndTime string `json:"purchase_price_end_time,omitempty" xml:"purchase_price_end_time,omitempty"`
	// 促销售价
	PromotionPrice string `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
	// 促销售价开始时间
	PromotionPriceStartTime string `json:"promotion_price_start_time,omitempty" xml:"promotion_price_start_time,omitempty"`
	// 促销售价结束时间
	PromotionPriceEndTime string `json:"promotion_price_end_time,omitempty" xml:"promotion_price_end_time,omitempty"`
	// 时间戳
	TimeStamp int64 `json:"time_stamp,omitempty" xml:"time_stamp,omitempty"`
}

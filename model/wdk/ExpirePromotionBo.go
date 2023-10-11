package wdk

// ExpirePromotionBo 结构体
type ExpirePromotionBo struct {
	// 短保时间段信息
	PeriodInfos []ExpirePeriodInfo `json:"period_infos,omitempty" xml:"period_infos>expire_period_info,omitempty"`
	// 门店id
	ShopIds []string `json:"shop_ids,omitempty" xml:"shop_ids>string,omitempty"`
	// 商品skucode
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 商家code
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
}

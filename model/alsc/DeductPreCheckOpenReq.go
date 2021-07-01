package alsc

// DeductPreCheckOpenReq 结构体
type DeductPreCheckOpenReq struct {
	// 品牌ID(不能和outbrandid同时为空))
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 卡Id，礼品卡或会员卡Id
	CardId string `json:"card_id,omitempty" xml:"card_id,omitempty"`
	// 会员
	CustomerId string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
	// 门店ID(不能和outshopid同时为空)
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 核销总资产
	Value int64 `json:"value,omitempty" xml:"value,omitempty"`
	// 外部门店ID(和shopid不能同时为空)
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 外部品牌ID(不能和brandid同时为空)
	OutBrandId string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
}

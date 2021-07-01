package alsc

// UnchargeCheckOpenReq 结构体
type UnchargeCheckOpenReq struct {
	// 品牌Id(不能和outbrandid同时为空)
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 卡号
	CardId string `json:"card_id,omitempty" xml:"card_id,omitempty"`
	// 顾客Id
	CustomerId string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
	// 门店id(不能和outshopid同时为空)
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 原充值订单Id
	SrcOuterOrderId string `json:"src_outer_order_id,omitempty" xml:"src_outer_order_id,omitempty"`
	// 外部门店ID(不能和shopid同时为空)
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 外部品牌ID(不能和brandid同时为空)
	OutBrandId string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
	// CS是辰森，KRY是客如云
	BizChannel string `json:"biz_channel,omitempty" xml:"biz_channel,omitempty"`
}

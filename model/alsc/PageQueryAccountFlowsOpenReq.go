package alsc

// PageQueryAccountFlowsOpenReq 结构体
type PageQueryAccountFlowsOpenReq struct {
	// 品牌ID(不能和outbrandid同时为空)
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 卡id
	CardId string `json:"card_id,omitempty" xml:"card_id,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 流水类型
	FlowTypes string `json:"flow_types,omitempty" xml:"flow_types,omitempty"`
	// 第几页,从1开始计数
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 每页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 门店ID(不能和outshopid同时为空)
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 会员id
	CustomerId string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
	// 外部门店ID(不能和shopid同时为空)
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 外部品牌ID(不能和brandid同时为空)
	OutBrandId string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
	// 外部订单ID
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// CS是辰森，KRY是客如云
	BizChannel string `json:"biz_channel,omitempty" xml:"biz_channel,omitempty"`
}

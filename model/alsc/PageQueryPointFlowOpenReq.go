package alsc

// PageQueryPointFlowOpenReq 结构体
type PageQueryPointFlowOpenReq struct {
	// 品牌id
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 顾客id
	CustomerId string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
	// 截止时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 关联交易ID,模糊查询
	OutBizId string `json:"out_biz_id,omitempty" xml:"out_biz_id,omitempty"`
	// 第几页,从1开始计数
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 每页大小，默认20
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 积分流水类型,类型参考枚举类
	PointFlowType string `json:"point_flow_type,omitempty" xml:"point_flow_type,omitempty"`
	// 门店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 外部门店ID,shop_id和out_shop_id不可同时为空
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 外部品牌id,brandId与out_brand_id不可同时为空
	OutBrandId string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
	// CS是辰森，KRY是客如云
	BizChannel string `json:"biz_channel,omitempty" xml:"biz_channel,omitempty"`
}

package alsc

// SearchCardOpenReq 结构体
type SearchCardOpenReq struct {
	// 品牌id
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 卡模板ID
	CardTemplateIds []string `json:"card_template_ids,omitempty" xml:"card_template_ids>string,omitempty"`
	// 卡类型
	CardType string `json:"card_type,omitempty" xml:"card_type,omitempty"`
	// 是否需要总数
	NeedCount bool `json:"need_count,omitempty" xml:"need_count,omitempty"`
	// 号码（手机号/卡实例ID/物理卡号）
	No string `json:"no,omitempty" xml:"no,omitempty"`
	// 开卡门店ID
	OpenCardShopIds []string `json:"open_card_shop_ids,omitempty" xml:"open_card_shop_ids>string,omitempty"`
	// 开卡时间终止
	OpenTimeEnd string `json:"open_time_end,omitempty" xml:"open_time_end,omitempty"`
	// 开卡时间开始
	OpenTimeStart string `json:"open_time_start,omitempty" xml:"open_time_start,omitempty"`
	// 外部品牌id,brandId与out_brand_id不可同时为空
	OutBrandId string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
	// 外部门店ID,shop_id和out_shop_id不可同时为空
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 第几页,从1开始计数
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 每页大小，默认20
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 门店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 卡实例状态
	Statuses []string `json:"statuses,omitempty" xml:"statuses>string,omitempty"`
}

package dt

// RtItemResearchPriceParam 结构体
type RtItemResearchPriceParam struct {
	// 流水号
	SerialNo string `json:"serial_no,omitempty" xml:"serial_no,omitempty"`
	// 业务来源
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 扩展字段
	ExtMap string `json:"ext_map,omitempty" xml:"ext_map,omitempty"`
	// 门店名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 区域名称
	AreaName string `json:"area_name,omitempty" xml:"area_name,omitempty"`
	// 原产地
	OriginPlace string `json:"origin_place,omitempty" xml:"origin_place,omitempty"`
	// 优惠结束时间
	PromotionEndTime string `json:"promotion_end_time,omitempty" xml:"promotion_end_time,omitempty"`
	// 优惠开始时间
	PromotionStartTime string `json:"promotion_start_time,omitempty" xml:"promotion_start_time,omitempty"`
	// 规格描述
	SpecDesc string `json:"spec_desc,omitempty" xml:"spec_desc,omitempty"`
	// 优惠描述
	PromotionDesc string `json:"promotion_desc,omitempty" xml:"promotion_desc,omitempty"`
	// 条码
	BarCode string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
	// 品牌
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 品名
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// rt货号
	RtItemNo string `json:"rt_item_no,omitempty" xml:"rt_item_no,omitempty"`
	// 门店Id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 区域id
	AreaId int64 `json:"area_id,omitempty" xml:"area_id,omitempty"`
	// 优惠类型
	PromotionType int64 `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
	// 是否促销
	InPromotion int64 `json:"in_promotion,omitempty" xml:"in_promotion,omitempty"`
	// 促销价
	PromotionPrice int64 `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
	// 正常价格
	SalePrice int64 `json:"sale_price,omitempty" xml:"sale_price,omitempty"`
}

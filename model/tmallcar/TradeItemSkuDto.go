package tmallcar

// TradeItemSkuDto 结构体
type TradeItemSkuDto struct {
	// 商家外部编码
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// sku名称等信息
	SkuInfo string `json:"sku_info,omitempty" xml:"sku_info,omitempty"`
	// skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

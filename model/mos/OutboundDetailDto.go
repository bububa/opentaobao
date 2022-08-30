package mos

// OutboundDetailDto 结构体
type OutboundDetailDto struct {
	// skuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 数量
	Quantity *BigDecimal `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

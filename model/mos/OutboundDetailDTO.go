package mos

// OutboundDetailDto 结构体
type OutboundDetailDto struct {
	// 数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// skuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

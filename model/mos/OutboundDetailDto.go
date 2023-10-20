package mos

import (
	"sync"
)

// OutboundDetailDto 结构体
type OutboundDetailDto struct {
	// skuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 数量
	Quantity float64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolOutboundDetailDto = sync.Pool{
	New: func() any {
		return new(OutboundDetailDto)
	},
}

// GetOutboundDetailDto() 从对象池中获取OutboundDetailDto
func GetOutboundDetailDto() *OutboundDetailDto {
	return poolOutboundDetailDto.Get().(*OutboundDetailDto)
}

// ReleaseOutboundDetailDto 释放OutboundDetailDto
func ReleaseOutboundDetailDto(v *OutboundDetailDto) {
	v.SkuId = ""
	v.Quantity = 0
	poolOutboundDetailDto.Put(v)
}

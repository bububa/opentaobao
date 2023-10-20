package pur

import (
	"sync"
)

// AccessLadderPriceDto 结构体
type AccessLadderPriceDto struct {
	// 原价
	OriginUnitPrice float64 `json:"origin_unit_price,omitempty" xml:"origin_unit_price,omitempty"`
	// 协议价
	UnitPrice float64 `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
	// 最小采购量
	MinimumPurchaseQuantity float64 `json:"minimum_purchase_quantity,omitempty" xml:"minimum_purchase_quantity,omitempty"`
}

var poolAccessLadderPriceDto = sync.Pool{
	New: func() any {
		return new(AccessLadderPriceDto)
	},
}

// GetAccessLadderPriceDto() 从对象池中获取AccessLadderPriceDto
func GetAccessLadderPriceDto() *AccessLadderPriceDto {
	return poolAccessLadderPriceDto.Get().(*AccessLadderPriceDto)
}

// ReleaseAccessLadderPriceDto 释放AccessLadderPriceDto
func ReleaseAccessLadderPriceDto(v *AccessLadderPriceDto) {
	v.OriginUnitPrice = 0
	v.UnitPrice = 0
	v.MinimumPurchaseQuantity = 0
	poolAccessLadderPriceDto.Put(v)
}

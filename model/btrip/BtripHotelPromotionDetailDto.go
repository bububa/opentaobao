package btrip

import (
	"sync"
)

// BtripHotelPromotionDetailDto 结构体
type BtripHotelPromotionDetailDto struct {
	// 优惠项名称
	PromotionName string `json:"promotion_name,omitempty" xml:"promotion_name,omitempty"`
	// 优惠金额
	PromotionPrice int64 `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
	// 优惠类型
	PromotionType int64 `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
}

var poolBtripHotelPromotionDetailDto = sync.Pool{
	New: func() any {
		return new(BtripHotelPromotionDetailDto)
	},
}

// GetBtripHotelPromotionDetailDto() 从对象池中获取BtripHotelPromotionDetailDto
func GetBtripHotelPromotionDetailDto() *BtripHotelPromotionDetailDto {
	return poolBtripHotelPromotionDetailDto.Get().(*BtripHotelPromotionDetailDto)
}

// ReleaseBtripHotelPromotionDetailDto 释放BtripHotelPromotionDetailDto
func ReleaseBtripHotelPromotionDetailDto(v *BtripHotelPromotionDetailDto) {
	v.PromotionName = ""
	v.PromotionPrice = 0
	v.PromotionType = 0
	poolBtripHotelPromotionDetailDto.Put(v)
}

package drug

import (
	"sync"
)

// OrderPromotionDto 结构体
type OrderPromotionDto struct {
	// 优惠名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 优惠类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 优惠金额
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

var poolOrderPromotionDto = sync.Pool{
	New: func() any {
		return new(OrderPromotionDto)
	},
}

// GetOrderPromotionDto() 从对象池中获取OrderPromotionDto
func GetOrderPromotionDto() *OrderPromotionDto {
	return poolOrderPromotionDto.Get().(*OrderPromotionDto)
}

// ReleaseOrderPromotionDto 释放OrderPromotionDto
func ReleaseOrderPromotionDto(v *OrderPromotionDto) {
	v.Name = ""
	v.Type = 0
	v.Amount = 0
	poolOrderPromotionDto.Put(v)
}

package logistic

import (
	"sync"
)

// DeliveryTimingDto 结构体
type DeliveryTimingDto struct {
	// 预估时效，这部分ISV直接展示，不要做改动，ISP会变文案和时效展示
	DeliveryPeriod string `json:"delivery_period,omitempty" xml:"delivery_period,omitempty"`
}

var poolDeliveryTimingDto = sync.Pool{
	New: func() any {
		return new(DeliveryTimingDto)
	},
}

// GetDeliveryTimingDto() 从对象池中获取DeliveryTimingDto
func GetDeliveryTimingDto() *DeliveryTimingDto {
	return poolDeliveryTimingDto.Get().(*DeliveryTimingDto)
}

// ReleaseDeliveryTimingDto 释放DeliveryTimingDto
func ReleaseDeliveryTimingDto(v *DeliveryTimingDto) {
	v.DeliveryPeriod = ""
	poolDeliveryTimingDto.Put(v)
}

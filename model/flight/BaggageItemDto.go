package flight

import (
	"sync"
)

// BaggageItemDto 结构体
type BaggageItemDto struct {
	// 行李业务对象集合
	BaggageBOList []BaggageDto `json:"baggage_b_o_list,omitempty" xml:"baggage_b_o_list>baggage_dto,omitempty"`
}

var poolBaggageItemDto = sync.Pool{
	New: func() any {
		return new(BaggageItemDto)
	},
}

// GetBaggageItemDto() 从对象池中获取BaggageItemDto
func GetBaggageItemDto() *BaggageItemDto {
	return poolBaggageItemDto.Get().(*BaggageItemDto)
}

// ReleaseBaggageItemDto 释放BaggageItemDto
func ReleaseBaggageItemDto(v *BaggageItemDto) {
	v.BaggageBOList = v.BaggageBOList[:0]
	poolBaggageItemDto.Put(v)
}

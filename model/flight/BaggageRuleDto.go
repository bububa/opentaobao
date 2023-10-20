package flight

import (
	"sync"
)

// BaggageRuleDto 结构体
type BaggageRuleDto struct {
	// 行李单元
	BaggageItem *BaggageItemDto `json:"baggage_item,omitempty" xml:"baggage_item,omitempty"`
}

var poolBaggageRuleDto = sync.Pool{
	New: func() any {
		return new(BaggageRuleDto)
	},
}

// GetBaggageRuleDto() 从对象池中获取BaggageRuleDto
func GetBaggageRuleDto() *BaggageRuleDto {
	return poolBaggageRuleDto.Get().(*BaggageRuleDto)
}

// ReleaseBaggageRuleDto 释放BaggageRuleDto
func ReleaseBaggageRuleDto(v *BaggageRuleDto) {
	v.BaggageItem = nil
	poolBaggageRuleDto.Put(v)
}

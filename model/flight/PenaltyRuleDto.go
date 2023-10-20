package flight

import (
	"sync"
)

// PenaltyRuleDto 结构体
type PenaltyRuleDto struct {
	// 退改签规则单元
	PenaltyItem *PenaltyItemDto `json:"penalty_item,omitempty" xml:"penalty_item,omitempty"`
}

var poolPenaltyRuleDto = sync.Pool{
	New: func() any {
		return new(PenaltyRuleDto)
	},
}

// GetPenaltyRuleDto() 从对象池中获取PenaltyRuleDto
func GetPenaltyRuleDto() *PenaltyRuleDto {
	return poolPenaltyRuleDto.Get().(*PenaltyRuleDto)
}

// ReleasePenaltyRuleDto 释放PenaltyRuleDto
func ReleasePenaltyRuleDto(v *PenaltyRuleDto) {
	v.PenaltyItem = nil
	poolPenaltyRuleDto.Put(v)
}

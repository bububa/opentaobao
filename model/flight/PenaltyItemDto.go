package flight

import (
	"sync"
)

// PenaltyItemDto 结构体
type PenaltyItemDto struct {
	// 退改签规则集合
	PenaltyBOList []PenaltyDto `json:"penalty_b_o_list,omitempty" xml:"penalty_b_o_list>penalty_dto,omitempty"`
}

var poolPenaltyItemDto = sync.Pool{
	New: func() any {
		return new(PenaltyItemDto)
	},
}

// GetPenaltyItemDto() 从对象池中获取PenaltyItemDto
func GetPenaltyItemDto() *PenaltyItemDto {
	return poolPenaltyItemDto.Get().(*PenaltyItemDto)
}

// ReleasePenaltyItemDto 释放PenaltyItemDto
func ReleasePenaltyItemDto(v *PenaltyItemDto) {
	v.PenaltyBOList = v.PenaltyBOList[:0]
	poolPenaltyItemDto.Put(v)
}

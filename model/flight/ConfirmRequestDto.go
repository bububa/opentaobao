package flight

import (
	"sync"
)

// ConfirmRequestDto 结构体
type ConfirmRequestDto struct {
	// 出价，要么为null，要么集合内必须有值
	Prices []float64 `json:"prices,omitempty" xml:"prices>float64,omitempty"`
	// 商家店铺id
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 意向单id
	IntentionId int64 `json:"intention_id,omitempty" xml:"intention_id,omitempty"`
}

var poolConfirmRequestDto = sync.Pool{
	New: func() any {
		return new(ConfirmRequestDto)
	},
}

// GetConfirmRequestDto() 从对象池中获取ConfirmRequestDto
func GetConfirmRequestDto() *ConfirmRequestDto {
	return poolConfirmRequestDto.Get().(*ConfirmRequestDto)
}

// ReleaseConfirmRequestDto 释放ConfirmRequestDto
func ReleaseConfirmRequestDto(v *ConfirmRequestDto) {
	v.Prices = v.Prices[:0]
	v.AgentId = 0
	v.IntentionId = 0
	poolConfirmRequestDto.Put(v)
}

package feedflow

import (
	"sync"
)

// IntelligentBidDto 结构体
type IntelligentBidDto struct {
	// 溢价范围
	ScopePercent int64 `json:"scope_percent,omitempty" xml:"scope_percent,omitempty"`
	// 人群策略 1:促进进店 2:促进成交
	Strategy int64 `json:"strategy,omitempty" xml:"strategy,omitempty"`
	// 是否打开
	Open bool `json:"open,omitempty" xml:"open,omitempty"`
}

var poolIntelligentBidDto = sync.Pool{
	New: func() any {
		return new(IntelligentBidDto)
	},
}

// GetIntelligentBidDto() 从对象池中获取IntelligentBidDto
func GetIntelligentBidDto() *IntelligentBidDto {
	return poolIntelligentBidDto.Get().(*IntelligentBidDto)
}

// ReleaseIntelligentBidDto 释放IntelligentBidDto
func ReleaseIntelligentBidDto(v *IntelligentBidDto) {
	v.ScopePercent = 0
	v.Strategy = 0
	v.Open = false
	poolIntelligentBidDto.Put(v)
}

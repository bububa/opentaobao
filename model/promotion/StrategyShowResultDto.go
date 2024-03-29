package promotion

import (
	"sync"
)

// StrategyShowResultDto 结构体
type StrategyShowResultDto struct {
	// 权益列表
	ShowBenefits []ShowBenefitDto `json:"show_benefits,omitempty" xml:"show_benefits>show_benefit_dto,omitempty"`
	// 扩展参数
	ExtraData string `json:"extra_data,omitempty" xml:"extra_data,omitempty"`
	// 追踪信息
	TrackingData string `json:"tracking_data,omitempty" xml:"tracking_data,omitempty"`
	// 投放计划
	ShowStrategy *ShowStrategyDto `json:"show_strategy,omitempty" xml:"show_strategy,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 是否有下一页
	HasNextPage bool `json:"has_next_page,omitempty" xml:"has_next_page,omitempty"`
}

var poolStrategyShowResultDto = sync.Pool{
	New: func() any {
		return new(StrategyShowResultDto)
	},
}

// GetStrategyShowResultDto() 从对象池中获取StrategyShowResultDto
func GetStrategyShowResultDto() *StrategyShowResultDto {
	return poolStrategyShowResultDto.Get().(*StrategyShowResultDto)
}

// ReleaseStrategyShowResultDto 释放StrategyShowResultDto
func ReleaseStrategyShowResultDto(v *StrategyShowResultDto) {
	v.ShowBenefits = v.ShowBenefits[:0]
	v.ExtraData = ""
	v.TrackingData = ""
	v.ShowStrategy = nil
	v.CurrentPage = 0
	v.HasNextPage = false
	poolStrategyShowResultDto.Put(v)
}

package scs

import (
	"sync"
)

// SolutionTopDto 结构体
type SolutionTopDto struct {
	// 计划信息
	CampaignView *CampaignSolutionTopDto `json:"campaign_view,omitempty" xml:"campaign_view,omitempty"`
}

var poolSolutionTopDto = sync.Pool{
	New: func() any {
		return new(SolutionTopDto)
	},
}

// GetSolutionTopDto() 从对象池中获取SolutionTopDto
func GetSolutionTopDto() *SolutionTopDto {
	return poolSolutionTopDto.Get().(*SolutionTopDto)
}

// ReleaseSolutionTopDto 释放SolutionTopDto
func ReleaseSolutionTopDto(v *SolutionTopDto) {
	v.CampaignView = nil
	poolSolutionTopDto.Put(v)
}

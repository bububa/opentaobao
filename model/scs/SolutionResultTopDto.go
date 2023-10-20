package scs

import (
	"sync"
)

// SolutionResultTopDto 结构体
type SolutionResultTopDto struct {
	// campaignId
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}

var poolSolutionResultTopDto = sync.Pool{
	New: func() any {
		return new(SolutionResultTopDto)
	},
}

// GetSolutionResultTopDto() 从对象池中获取SolutionResultTopDto
func GetSolutionResultTopDto() *SolutionResultTopDto {
	return poolSolutionResultTopDto.Get().(*SolutionResultTopDto)
}

// ReleaseSolutionResultTopDto 释放SolutionResultTopDto
func ReleaseSolutionResultTopDto(v *SolutionResultTopDto) {
	v.CampaignId = 0
	poolSolutionResultTopDto.Put(v)
}

package feedflow

import (
	"sync"
)

// CrowdQueryDto 结构体
type CrowdQueryDto struct {
	// 定向类型
	TargetTypeList []string `json:"target_type_list,omitempty" xml:"target_type_list>string,omitempty"`
	// 人群状态
	StatusList []string `json:"status_list,omitempty" xml:"status_list>string,omitempty"`
	// 单元id
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
	// 分页条件
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 人群id
	CrowdId int64 `json:"crowd_id,omitempty" xml:"crowd_id,omitempty"`
	// 分页条件
	Offset int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}

var poolCrowdQueryDto = sync.Pool{
	New: func() any {
		return new(CrowdQueryDto)
	},
}

// GetCrowdQueryDto() 从对象池中获取CrowdQueryDto
func GetCrowdQueryDto() *CrowdQueryDto {
	return poolCrowdQueryDto.Get().(*CrowdQueryDto)
}

// ReleaseCrowdQueryDto 释放CrowdQueryDto
func ReleaseCrowdQueryDto(v *CrowdQueryDto) {
	v.TargetTypeList = v.TargetTypeList[:0]
	v.StatusList = v.StatusList[:0]
	v.AdgroupId = 0
	v.PageSize = 0
	v.CrowdId = 0
	v.Offset = 0
	v.CampaignId = 0
	poolCrowdQueryDto.Put(v)
}

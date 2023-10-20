package feedflow

import (
	"sync"
)

// CampaignQueryDto 结构体
type CampaignQueryDto struct {
	// 状态列表
	StatusList []string `json:"status_list,omitempty" xml:"status_list>string,omitempty"`
	// 计划名称
	CampaignName string `json:"campaign_name,omitempty" xml:"campaign_name,omitempty"`
	// 计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 起始位置
	Offset int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 每页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolCampaignQueryDto = sync.Pool{
	New: func() any {
		return new(CampaignQueryDto)
	},
}

// GetCampaignQueryDto() 从对象池中获取CampaignQueryDto
func GetCampaignQueryDto() *CampaignQueryDto {
	return poolCampaignQueryDto.Get().(*CampaignQueryDto)
}

// ReleaseCampaignQueryDto 释放CampaignQueryDto
func ReleaseCampaignQueryDto(v *CampaignQueryDto) {
	v.StatusList = v.StatusList[:0]
	v.CampaignName = ""
	v.CampaignId = 0
	v.Offset = 0
	v.PageSize = 0
	poolCampaignQueryDto.Put(v)
}

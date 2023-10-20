package scbp

import (
	"sync"
)

// TopP4pModifyQuickCampaignDto 结构体
type TopP4pModifyQuickCampaignDto struct {
	// 操作类型，0=计划暂停，1=计划开启，2=计划删除
	Action int64 `json:"action,omitempty" xml:"action,omitempty"`
	// 计划ID
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}

var poolTopP4pModifyQuickCampaignDto = sync.Pool{
	New: func() any {
		return new(TopP4pModifyQuickCampaignDto)
	},
}

// GetTopP4pModifyQuickCampaignDto() 从对象池中获取TopP4pModifyQuickCampaignDto
func GetTopP4pModifyQuickCampaignDto() *TopP4pModifyQuickCampaignDto {
	return poolTopP4pModifyQuickCampaignDto.Get().(*TopP4pModifyQuickCampaignDto)
}

// ReleaseTopP4pModifyQuickCampaignDto 释放TopP4pModifyQuickCampaignDto
func ReleaseTopP4pModifyQuickCampaignDto(v *TopP4pModifyQuickCampaignDto) {
	v.Action = 0
	v.CampaignId = 0
	poolTopP4pModifyQuickCampaignDto.Put(v)
}

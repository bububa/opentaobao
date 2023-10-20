package scbp

import (
	"sync"
)

// CampaignTargetingWordSettingOperationDto 结构体
type CampaignTargetingWordSettingOperationDto struct {
	// 操作优推类型 add-增 del-删 mod-改
	Operation string `json:"operation,omitempty" xml:"operation,omitempty"`
	// 关键词
	Keyword string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 待操作得优推品id
	AdgroupId string `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
	// 关键词id
	KeywordId int64 `json:"keyword_id,omitempty" xml:"keyword_id,omitempty"`
	// 计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}

var poolCampaignTargetingWordSettingOperationDto = sync.Pool{
	New: func() any {
		return new(CampaignTargetingWordSettingOperationDto)
	},
}

// GetCampaignTargetingWordSettingOperationDto() 从对象池中获取CampaignTargetingWordSettingOperationDto
func GetCampaignTargetingWordSettingOperationDto() *CampaignTargetingWordSettingOperationDto {
	return poolCampaignTargetingWordSettingOperationDto.Get().(*CampaignTargetingWordSettingOperationDto)
}

// ReleaseCampaignTargetingWordSettingOperationDto 释放CampaignTargetingWordSettingOperationDto
func ReleaseCampaignTargetingWordSettingOperationDto(v *CampaignTargetingWordSettingOperationDto) {
	v.Operation = ""
	v.Keyword = ""
	v.AdgroupId = ""
	v.KeywordId = 0
	v.CampaignId = 0
	poolCampaignTargetingWordSettingOperationDto.Put(v)
}

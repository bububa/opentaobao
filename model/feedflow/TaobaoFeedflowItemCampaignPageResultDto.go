package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemCampaignPageResultDto 结构体
type TaobaoFeedflowItemCampaignPageResultDto struct {
	// 计划列表
	Results []CampaignDto `json:"results,omitempty" xml:"results>campaign_dto,omitempty"`
	// 信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 符合条件的计划数量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemCampaignPageResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCampaignPageResultDto)
	},
}

// GetTaobaoFeedflowItemCampaignPageResultDto() 从对象池中获取TaobaoFeedflowItemCampaignPageResultDto
func GetTaobaoFeedflowItemCampaignPageResultDto() *TaobaoFeedflowItemCampaignPageResultDto {
	return poolTaobaoFeedflowItemCampaignPageResultDto.Get().(*TaobaoFeedflowItemCampaignPageResultDto)
}

// ReleaseTaobaoFeedflowItemCampaignPageResultDto 释放TaobaoFeedflowItemCampaignPageResultDto
func ReleaseTaobaoFeedflowItemCampaignPageResultDto(v *TaobaoFeedflowItemCampaignPageResultDto) {
	v.Results = v.Results[:0]
	v.Message = ""
	v.TotalCount = 0
	v.Success = false
	poolTaobaoFeedflowItemCampaignPageResultDto.Put(v)
}

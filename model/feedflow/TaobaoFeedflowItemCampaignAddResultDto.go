package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemCampaignAddResultDto 结构体
type TaobaoFeedflowItemCampaignAddResultDto struct {
	// 系统自动生成
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 结果
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
	// 系统自动生成
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemCampaignAddResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCampaignAddResultDto)
	},
}

// GetTaobaoFeedflowItemCampaignAddResultDto() 从对象池中获取TaobaoFeedflowItemCampaignAddResultDto
func GetTaobaoFeedflowItemCampaignAddResultDto() *TaobaoFeedflowItemCampaignAddResultDto {
	return poolTaobaoFeedflowItemCampaignAddResultDto.Get().(*TaobaoFeedflowItemCampaignAddResultDto)
}

// ReleaseTaobaoFeedflowItemCampaignAddResultDto 释放TaobaoFeedflowItemCampaignAddResultDto
func ReleaseTaobaoFeedflowItemCampaignAddResultDto(v *TaobaoFeedflowItemCampaignAddResultDto) {
	v.Message = ""
	v.Result = 0
	v.Success = false
	poolTaobaoFeedflowItemCampaignAddResultDto.Put(v)
}

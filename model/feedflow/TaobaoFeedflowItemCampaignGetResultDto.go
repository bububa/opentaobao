package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemCampaignGetResultDto 结构体
type TaobaoFeedflowItemCampaignGetResultDto struct {
	// 信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 计划信息
	Result *CampaignDto `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemCampaignGetResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCampaignGetResultDto)
	},
}

// GetTaobaoFeedflowItemCampaignGetResultDto() 从对象池中获取TaobaoFeedflowItemCampaignGetResultDto
func GetTaobaoFeedflowItemCampaignGetResultDto() *TaobaoFeedflowItemCampaignGetResultDto {
	return poolTaobaoFeedflowItemCampaignGetResultDto.Get().(*TaobaoFeedflowItemCampaignGetResultDto)
}

// ReleaseTaobaoFeedflowItemCampaignGetResultDto 释放TaobaoFeedflowItemCampaignGetResultDto
func ReleaseTaobaoFeedflowItemCampaignGetResultDto(v *TaobaoFeedflowItemCampaignGetResultDto) {
	v.Message = ""
	v.Result = nil
	v.Success = false
	poolTaobaoFeedflowItemCampaignGetResultDto.Put(v)
}

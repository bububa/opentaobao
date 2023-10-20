package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemCampaignModifyResultDto 结构体
type TaobaoFeedflowItemCampaignModifyResultDto struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemCampaignModifyResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCampaignModifyResultDto)
	},
}

// GetTaobaoFeedflowItemCampaignModifyResultDto() 从对象池中获取TaobaoFeedflowItemCampaignModifyResultDto
func GetTaobaoFeedflowItemCampaignModifyResultDto() *TaobaoFeedflowItemCampaignModifyResultDto {
	return poolTaobaoFeedflowItemCampaignModifyResultDto.Get().(*TaobaoFeedflowItemCampaignModifyResultDto)
}

// ReleaseTaobaoFeedflowItemCampaignModifyResultDto 释放TaobaoFeedflowItemCampaignModifyResultDto
func ReleaseTaobaoFeedflowItemCampaignModifyResultDto(v *TaobaoFeedflowItemCampaignModifyResultDto) {
	v.Message = ""
	v.Result = false
	v.Success = false
	poolTaobaoFeedflowItemCampaignModifyResultDto.Put(v)
}

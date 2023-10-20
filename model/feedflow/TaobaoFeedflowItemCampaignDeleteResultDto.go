package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemCampaignDeleteResultDto 结构体
type TaobaoFeedflowItemCampaignDeleteResultDto struct {
	// 信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 操作结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemCampaignDeleteResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCampaignDeleteResultDto)
	},
}

// GetTaobaoFeedflowItemCampaignDeleteResultDto() 从对象池中获取TaobaoFeedflowItemCampaignDeleteResultDto
func GetTaobaoFeedflowItemCampaignDeleteResultDto() *TaobaoFeedflowItemCampaignDeleteResultDto {
	return poolTaobaoFeedflowItemCampaignDeleteResultDto.Get().(*TaobaoFeedflowItemCampaignDeleteResultDto)
}

// ReleaseTaobaoFeedflowItemCampaignDeleteResultDto 释放TaobaoFeedflowItemCampaignDeleteResultDto
func ReleaseTaobaoFeedflowItemCampaignDeleteResultDto(v *TaobaoFeedflowItemCampaignDeleteResultDto) {
	v.Message = ""
	v.Result = false
	v.Success = false
	poolTaobaoFeedflowItemCampaignDeleteResultDto.Put(v)
}

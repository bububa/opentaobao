package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemCampaignDaybudgetResultDto 结构体
type TaobaoFeedflowItemCampaignDaybudgetResultDto struct {
	// 信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 预算总额
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemCampaignDaybudgetResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCampaignDaybudgetResultDto)
	},
}

// GetTaobaoFeedflowItemCampaignDaybudgetResultDto() 从对象池中获取TaobaoFeedflowItemCampaignDaybudgetResultDto
func GetTaobaoFeedflowItemCampaignDaybudgetResultDto() *TaobaoFeedflowItemCampaignDaybudgetResultDto {
	return poolTaobaoFeedflowItemCampaignDaybudgetResultDto.Get().(*TaobaoFeedflowItemCampaignDaybudgetResultDto)
}

// ReleaseTaobaoFeedflowItemCampaignDaybudgetResultDto 释放TaobaoFeedflowItemCampaignDaybudgetResultDto
func ReleaseTaobaoFeedflowItemCampaignDaybudgetResultDto(v *TaobaoFeedflowItemCampaignDaybudgetResultDto) {
	v.Message = ""
	v.Result = 0
	v.Success = false
	poolTaobaoFeedflowItemCampaignDaybudgetResultDto.Put(v)
}

package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemCampaignRptdailylistResultDto 结构体
type TaobaoFeedflowItemCampaignRptdailylistResultDto struct {
	// 报表信息
	RptList []RptResultDto `json:"rpt_list,omitempty" xml:"rpt_list>rpt_result_dto,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemCampaignRptdailylistResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCampaignRptdailylistResultDto)
	},
}

// GetTaobaoFeedflowItemCampaignRptdailylistResultDto() 从对象池中获取TaobaoFeedflowItemCampaignRptdailylistResultDto
func GetTaobaoFeedflowItemCampaignRptdailylistResultDto() *TaobaoFeedflowItemCampaignRptdailylistResultDto {
	return poolTaobaoFeedflowItemCampaignRptdailylistResultDto.Get().(*TaobaoFeedflowItemCampaignRptdailylistResultDto)
}

// ReleaseTaobaoFeedflowItemCampaignRptdailylistResultDto 释放TaobaoFeedflowItemCampaignRptdailylistResultDto
func ReleaseTaobaoFeedflowItemCampaignRptdailylistResultDto(v *TaobaoFeedflowItemCampaignRptdailylistResultDto) {
	v.RptList = v.RptList[:0]
	v.Message = ""
	v.Success = false
	poolTaobaoFeedflowItemCampaignRptdailylistResultDto.Put(v)
}

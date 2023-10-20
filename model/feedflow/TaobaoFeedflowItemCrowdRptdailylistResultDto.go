package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemCrowdRptdailylistResultDto 结构体
type TaobaoFeedflowItemCrowdRptdailylistResultDto struct {
	// 报表信息
	RptList []RptResultDto `json:"rpt_list,omitempty" xml:"rpt_list>rpt_result_dto,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemCrowdRptdailylistResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCrowdRptdailylistResultDto)
	},
}

// GetTaobaoFeedflowItemCrowdRptdailylistResultDto() 从对象池中获取TaobaoFeedflowItemCrowdRptdailylistResultDto
func GetTaobaoFeedflowItemCrowdRptdailylistResultDto() *TaobaoFeedflowItemCrowdRptdailylistResultDto {
	return poolTaobaoFeedflowItemCrowdRptdailylistResultDto.Get().(*TaobaoFeedflowItemCrowdRptdailylistResultDto)
}

// ReleaseTaobaoFeedflowItemCrowdRptdailylistResultDto 释放TaobaoFeedflowItemCrowdRptdailylistResultDto
func ReleaseTaobaoFeedflowItemCrowdRptdailylistResultDto(v *TaobaoFeedflowItemCrowdRptdailylistResultDto) {
	v.RptList = v.RptList[:0]
	v.Message = ""
	v.Success = false
	poolTaobaoFeedflowItemCrowdRptdailylistResultDto.Put(v)
}

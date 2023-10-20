package feedflow

import (
	"sync"
)

// TaobaoFeedflowAccountRptdailylistResultDto 结构体
type TaobaoFeedflowAccountRptdailylistResultDto struct {
	// 报表信息
	RptList []RptResultDto `json:"rpt_list,omitempty" xml:"rpt_list>rpt_result_dto,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowAccountRptdailylistResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowAccountRptdailylistResultDto)
	},
}

// GetTaobaoFeedflowAccountRptdailylistResultDto() 从对象池中获取TaobaoFeedflowAccountRptdailylistResultDto
func GetTaobaoFeedflowAccountRptdailylistResultDto() *TaobaoFeedflowAccountRptdailylistResultDto {
	return poolTaobaoFeedflowAccountRptdailylistResultDto.Get().(*TaobaoFeedflowAccountRptdailylistResultDto)
}

// ReleaseTaobaoFeedflowAccountRptdailylistResultDto 释放TaobaoFeedflowAccountRptdailylistResultDto
func ReleaseTaobaoFeedflowAccountRptdailylistResultDto(v *TaobaoFeedflowAccountRptdailylistResultDto) {
	v.RptList = v.RptList[:0]
	v.Message = ""
	v.Success = false
	poolTaobaoFeedflowAccountRptdailylistResultDto.Put(v)
}

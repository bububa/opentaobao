package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemAdzoneRptdailylistResultDto 结构体
type TaobaoFeedflowItemAdzoneRptdailylistResultDto struct {
	// 报表信息
	RptList []RptResultDto `json:"rpt_list,omitempty" xml:"rpt_list>rpt_result_dto,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemAdzoneRptdailylistResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAdzoneRptdailylistResultDto)
	},
}

// GetTaobaoFeedflowItemAdzoneRptdailylistResultDto() 从对象池中获取TaobaoFeedflowItemAdzoneRptdailylistResultDto
func GetTaobaoFeedflowItemAdzoneRptdailylistResultDto() *TaobaoFeedflowItemAdzoneRptdailylistResultDto {
	return poolTaobaoFeedflowItemAdzoneRptdailylistResultDto.Get().(*TaobaoFeedflowItemAdzoneRptdailylistResultDto)
}

// ReleaseTaobaoFeedflowItemAdzoneRptdailylistResultDto 释放TaobaoFeedflowItemAdzoneRptdailylistResultDto
func ReleaseTaobaoFeedflowItemAdzoneRptdailylistResultDto(v *TaobaoFeedflowItemAdzoneRptdailylistResultDto) {
	v.RptList = v.RptList[:0]
	v.Message = ""
	v.Success = false
	poolTaobaoFeedflowItemAdzoneRptdailylistResultDto.Put(v)
}

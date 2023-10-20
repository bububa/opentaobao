package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemCreativeRptdailylistResultDto 结构体
type TaobaoFeedflowItemCreativeRptdailylistResultDto struct {
	// 报表结果
	RptList []RptResultDto `json:"rpt_list,omitempty" xml:"rpt_list>rpt_result_dto,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemCreativeRptdailylistResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCreativeRptdailylistResultDto)
	},
}

// GetTaobaoFeedflowItemCreativeRptdailylistResultDto() 从对象池中获取TaobaoFeedflowItemCreativeRptdailylistResultDto
func GetTaobaoFeedflowItemCreativeRptdailylistResultDto() *TaobaoFeedflowItemCreativeRptdailylistResultDto {
	return poolTaobaoFeedflowItemCreativeRptdailylistResultDto.Get().(*TaobaoFeedflowItemCreativeRptdailylistResultDto)
}

// ReleaseTaobaoFeedflowItemCreativeRptdailylistResultDto 释放TaobaoFeedflowItemCreativeRptdailylistResultDto
func ReleaseTaobaoFeedflowItemCreativeRptdailylistResultDto(v *TaobaoFeedflowItemCreativeRptdailylistResultDto) {
	v.RptList = v.RptList[:0]
	v.Message = ""
	v.Success = false
	poolTaobaoFeedflowItemCreativeRptdailylistResultDto.Put(v)
}

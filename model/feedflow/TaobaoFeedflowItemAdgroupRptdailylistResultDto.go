package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemAdgroupRptdailylistResultDto 结构体
type TaobaoFeedflowItemAdgroupRptdailylistResultDto struct {
	// 报表信息
	RptList []RptResultDto `json:"rpt_list,omitempty" xml:"rpt_list>rpt_result_dto,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemAdgroupRptdailylistResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAdgroupRptdailylistResultDto)
	},
}

// GetTaobaoFeedflowItemAdgroupRptdailylistResultDto() 从对象池中获取TaobaoFeedflowItemAdgroupRptdailylistResultDto
func GetTaobaoFeedflowItemAdgroupRptdailylistResultDto() *TaobaoFeedflowItemAdgroupRptdailylistResultDto {
	return poolTaobaoFeedflowItemAdgroupRptdailylistResultDto.Get().(*TaobaoFeedflowItemAdgroupRptdailylistResultDto)
}

// ReleaseTaobaoFeedflowItemAdgroupRptdailylistResultDto 释放TaobaoFeedflowItemAdgroupRptdailylistResultDto
func ReleaseTaobaoFeedflowItemAdgroupRptdailylistResultDto(v *TaobaoFeedflowItemAdgroupRptdailylistResultDto) {
	v.RptList = v.RptList[:0]
	v.Message = ""
	v.Success = false
	poolTaobaoFeedflowItemAdgroupRptdailylistResultDto.Put(v)
}

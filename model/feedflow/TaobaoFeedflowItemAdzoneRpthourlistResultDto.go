package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemAdzoneRpthourlistResultDto 结构体
type TaobaoFeedflowItemAdzoneRpthourlistResultDto struct {
	// 返回结果
	RptList []RptResultDto `json:"rpt_list,omitempty" xml:"rpt_list>rpt_result_dto,omitempty"`
	// 描述信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 返回信息
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemAdzoneRpthourlistResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAdzoneRpthourlistResultDto)
	},
}

// GetTaobaoFeedflowItemAdzoneRpthourlistResultDto() 从对象池中获取TaobaoFeedflowItemAdzoneRpthourlistResultDto
func GetTaobaoFeedflowItemAdzoneRpthourlistResultDto() *TaobaoFeedflowItemAdzoneRpthourlistResultDto {
	return poolTaobaoFeedflowItemAdzoneRpthourlistResultDto.Get().(*TaobaoFeedflowItemAdzoneRpthourlistResultDto)
}

// ReleaseTaobaoFeedflowItemAdzoneRpthourlistResultDto 释放TaobaoFeedflowItemAdzoneRpthourlistResultDto
func ReleaseTaobaoFeedflowItemAdzoneRpthourlistResultDto(v *TaobaoFeedflowItemAdzoneRpthourlistResultDto) {
	v.RptList = v.RptList[:0]
	v.Message = ""
	v.TotalCount = 0
	v.ResultCode = nil
	v.Success = false
	poolTaobaoFeedflowItemAdzoneRpthourlistResultDto.Put(v)
}

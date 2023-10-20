package feedflow

import (
	"sync"
)

// TaobaoFeedflowAccountRpthourlistResultDto 结构体
type TaobaoFeedflowAccountRpthourlistResultDto struct {
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

var poolTaobaoFeedflowAccountRpthourlistResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowAccountRpthourlistResultDto)
	},
}

// GetTaobaoFeedflowAccountRpthourlistResultDto() 从对象池中获取TaobaoFeedflowAccountRpthourlistResultDto
func GetTaobaoFeedflowAccountRpthourlistResultDto() *TaobaoFeedflowAccountRpthourlistResultDto {
	return poolTaobaoFeedflowAccountRpthourlistResultDto.Get().(*TaobaoFeedflowAccountRpthourlistResultDto)
}

// ReleaseTaobaoFeedflowAccountRpthourlistResultDto 释放TaobaoFeedflowAccountRpthourlistResultDto
func ReleaseTaobaoFeedflowAccountRpthourlistResultDto(v *TaobaoFeedflowAccountRpthourlistResultDto) {
	v.RptList = v.RptList[:0]
	v.Message = ""
	v.TotalCount = 0
	v.ResultCode = nil
	v.Success = false
	poolTaobaoFeedflowAccountRpthourlistResultDto.Put(v)
}

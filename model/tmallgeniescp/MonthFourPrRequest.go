package tmallgeniescp

import (
	"sync"
)

// MonthFourPrRequest 结构体
type MonthFourPrRequest struct {
	// 请求参数
	MonthFourPrParamDTOS []MonthFourPrParamDto `json:"month_four_pr_param_d_t_o_s,omitempty" xml:"month_four_pr_param_d_t_o_s>month_four_pr_param_dto,omitempty"`
	// 扩展参数
	RequestExtendJson string `json:"request_extend_json,omitempty" xml:"request_extend_json,omitempty"`
}

var poolMonthFourPrRequest = sync.Pool{
	New: func() any {
		return new(MonthFourPrRequest)
	},
}

// GetMonthFourPrRequest() 从对象池中获取MonthFourPrRequest
func GetMonthFourPrRequest() *MonthFourPrRequest {
	return poolMonthFourPrRequest.Get().(*MonthFourPrRequest)
}

// ReleaseMonthFourPrRequest 释放MonthFourPrRequest
func ReleaseMonthFourPrRequest(v *MonthFourPrRequest) {
	v.MonthFourPrParamDTOS = v.MonthFourPrParamDTOS[:0]
	v.RequestExtendJson = ""
	poolMonthFourPrRequest.Put(v)
}

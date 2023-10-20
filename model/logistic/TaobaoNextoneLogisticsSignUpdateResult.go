package logistic

import (
	"sync"
)

// TaobaoNextoneLogisticsSignUpdateResult 结构体
type TaobaoNextoneLogisticsSignUpdateResult struct {
	// 返回数据
	ResultData string `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// 错误信息
	ErrorInfo string `json:"error_info,omitempty" xml:"error_info,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 成功失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoNextoneLogisticsSignUpdateResult = sync.Pool{
	New: func() any {
		return new(TaobaoNextoneLogisticsSignUpdateResult)
	},
}

// GetTaobaoNextoneLogisticsSignUpdateResult() 从对象池中获取TaobaoNextoneLogisticsSignUpdateResult
func GetTaobaoNextoneLogisticsSignUpdateResult() *TaobaoNextoneLogisticsSignUpdateResult {
	return poolTaobaoNextoneLogisticsSignUpdateResult.Get().(*TaobaoNextoneLogisticsSignUpdateResult)
}

// ReleaseTaobaoNextoneLogisticsSignUpdateResult 释放TaobaoNextoneLogisticsSignUpdateResult
func ReleaseTaobaoNextoneLogisticsSignUpdateResult(v *TaobaoNextoneLogisticsSignUpdateResult) {
	v.ResultData = ""
	v.ErrorInfo = ""
	v.ErrorCode = ""
	v.Success = false
	poolTaobaoNextoneLogisticsSignUpdateResult.Put(v)
}

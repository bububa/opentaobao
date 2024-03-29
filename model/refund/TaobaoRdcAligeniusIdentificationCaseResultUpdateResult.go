package refund

import (
	"sync"
)

// TaobaoRdcAligeniusIdentificationCaseResultUpdateResult 结构体
type TaobaoRdcAligeniusIdentificationCaseResultUpdateResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorInfo string `json:"error_info,omitempty" xml:"error_info,omitempty"`
	// resultData
	ResultData *Resultdata `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoRdcAligeniusIdentificationCaseResultUpdateResult = sync.Pool{
	New: func() any {
		return new(TaobaoRdcAligeniusIdentificationCaseResultUpdateResult)
	},
}

// GetTaobaoRdcAligeniusIdentificationCaseResultUpdateResult() 从对象池中获取TaobaoRdcAligeniusIdentificationCaseResultUpdateResult
func GetTaobaoRdcAligeniusIdentificationCaseResultUpdateResult() *TaobaoRdcAligeniusIdentificationCaseResultUpdateResult {
	return poolTaobaoRdcAligeniusIdentificationCaseResultUpdateResult.Get().(*TaobaoRdcAligeniusIdentificationCaseResultUpdateResult)
}

// ReleaseTaobaoRdcAligeniusIdentificationCaseResultUpdateResult 释放TaobaoRdcAligeniusIdentificationCaseResultUpdateResult
func ReleaseTaobaoRdcAligeniusIdentificationCaseResultUpdateResult(v *TaobaoRdcAligeniusIdentificationCaseResultUpdateResult) {
	v.ErrorCode = ""
	v.ErrorInfo = ""
	v.ResultData = nil
	v.Success = false
	poolTaobaoRdcAligeniusIdentificationCaseResultUpdateResult.Put(v)
}

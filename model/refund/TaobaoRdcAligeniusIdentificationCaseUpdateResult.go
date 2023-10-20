package refund

import (
	"sync"
)

// TaobaoRdcAligeniusIdentificationCaseUpdateResult 结构体
type TaobaoRdcAligeniusIdentificationCaseUpdateResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorInfo string `json:"error_info,omitempty" xml:"error_info,omitempty"`
	// resultData
	ResultData *Resultdata `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoRdcAligeniusIdentificationCaseUpdateResult = sync.Pool{
	New: func() any {
		return new(TaobaoRdcAligeniusIdentificationCaseUpdateResult)
	},
}

// GetTaobaoRdcAligeniusIdentificationCaseUpdateResult() 从对象池中获取TaobaoRdcAligeniusIdentificationCaseUpdateResult
func GetTaobaoRdcAligeniusIdentificationCaseUpdateResult() *TaobaoRdcAligeniusIdentificationCaseUpdateResult {
	return poolTaobaoRdcAligeniusIdentificationCaseUpdateResult.Get().(*TaobaoRdcAligeniusIdentificationCaseUpdateResult)
}

// ReleaseTaobaoRdcAligeniusIdentificationCaseUpdateResult 释放TaobaoRdcAligeniusIdentificationCaseUpdateResult
func ReleaseTaobaoRdcAligeniusIdentificationCaseUpdateResult(v *TaobaoRdcAligeniusIdentificationCaseUpdateResult) {
	v.ErrorCode = ""
	v.ErrorInfo = ""
	v.ResultData = nil
	v.Success = false
	poolTaobaoRdcAligeniusIdentificationCaseUpdateResult.Put(v)
}

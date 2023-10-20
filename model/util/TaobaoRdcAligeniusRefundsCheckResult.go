package util

import (
	"sync"
)

// TaobaoRdcAligeniusRefundsCheckResult 结构体
type TaobaoRdcAligeniusRefundsCheckResult struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorInfo
	ErrorInfo string `json:"error_info,omitempty" xml:"error_info,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoRdcAligeniusRefundsCheckResult = sync.Pool{
	New: func() any {
		return new(TaobaoRdcAligeniusRefundsCheckResult)
	},
}

// GetTaobaoRdcAligeniusRefundsCheckResult() 从对象池中获取TaobaoRdcAligeniusRefundsCheckResult
func GetTaobaoRdcAligeniusRefundsCheckResult() *TaobaoRdcAligeniusRefundsCheckResult {
	return poolTaobaoRdcAligeniusRefundsCheckResult.Get().(*TaobaoRdcAligeniusRefundsCheckResult)
}

// ReleaseTaobaoRdcAligeniusRefundsCheckResult 释放TaobaoRdcAligeniusRefundsCheckResult
func ReleaseTaobaoRdcAligeniusRefundsCheckResult(v *TaobaoRdcAligeniusRefundsCheckResult) {
	v.ErrorCode = ""
	v.ErrorInfo = ""
	v.Success = false
	poolTaobaoRdcAligeniusRefundsCheckResult.Put(v)
}

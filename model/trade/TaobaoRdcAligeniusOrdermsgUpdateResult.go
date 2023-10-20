package trade

import (
	"sync"
)

// TaobaoRdcAligeniusOrdermsgUpdateResult 结构体
type TaobaoRdcAligeniusOrdermsgUpdateResult struct {
	// resultData
	ResultData string `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// errorInfo
	ErrorInfo string `json:"error_info,omitempty" xml:"error_info,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoRdcAligeniusOrdermsgUpdateResult = sync.Pool{
	New: func() any {
		return new(TaobaoRdcAligeniusOrdermsgUpdateResult)
	},
}

// GetTaobaoRdcAligeniusOrdermsgUpdateResult() 从对象池中获取TaobaoRdcAligeniusOrdermsgUpdateResult
func GetTaobaoRdcAligeniusOrdermsgUpdateResult() *TaobaoRdcAligeniusOrdermsgUpdateResult {
	return poolTaobaoRdcAligeniusOrdermsgUpdateResult.Get().(*TaobaoRdcAligeniusOrdermsgUpdateResult)
}

// ReleaseTaobaoRdcAligeniusOrdermsgUpdateResult 释放TaobaoRdcAligeniusOrdermsgUpdateResult
func ReleaseTaobaoRdcAligeniusOrdermsgUpdateResult(v *TaobaoRdcAligeniusOrdermsgUpdateResult) {
	v.ResultData = ""
	v.ErrorInfo = ""
	v.ErrorCode = ""
	v.Success = false
	poolTaobaoRdcAligeniusOrdermsgUpdateResult.Put(v)
}

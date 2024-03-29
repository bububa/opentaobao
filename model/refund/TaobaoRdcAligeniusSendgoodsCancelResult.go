package refund

import (
	"sync"
)

// TaobaoRdcAligeniusSendgoodsCancelResult 结构体
type TaobaoRdcAligeniusSendgoodsCancelResult struct {
	// 异常信息
	ErrorInfo string `json:"error_info,omitempty" xml:"error_info,omitempty"`
	// 异常编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// resultData
	ResultData *Resultdata `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoRdcAligeniusSendgoodsCancelResult = sync.Pool{
	New: func() any {
		return new(TaobaoRdcAligeniusSendgoodsCancelResult)
	},
}

// GetTaobaoRdcAligeniusSendgoodsCancelResult() 从对象池中获取TaobaoRdcAligeniusSendgoodsCancelResult
func GetTaobaoRdcAligeniusSendgoodsCancelResult() *TaobaoRdcAligeniusSendgoodsCancelResult {
	return poolTaobaoRdcAligeniusSendgoodsCancelResult.Get().(*TaobaoRdcAligeniusSendgoodsCancelResult)
}

// ReleaseTaobaoRdcAligeniusSendgoodsCancelResult 释放TaobaoRdcAligeniusSendgoodsCancelResult
func ReleaseTaobaoRdcAligeniusSendgoodsCancelResult(v *TaobaoRdcAligeniusSendgoodsCancelResult) {
	v.ErrorInfo = ""
	v.ErrorCode = ""
	v.ResultData = nil
	v.Success = false
	poolTaobaoRdcAligeniusSendgoodsCancelResult.Put(v)
}

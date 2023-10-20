package tmallservice

import (
	"sync"
)

// TmallServiceSettleadjustmentRequestResult 结构体
type TmallServiceSettleadjustmentRequestResult struct {
	// dataModule
	DataModule *SettleAdjustmentResp `json:"data_module,omitempty" xml:"data_module,omitempty"`
	// errorMessage
	ErrorMessage *ErrorMessage `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// true：查询成功，false：失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServiceSettleadjustmentRequestResult = sync.Pool{
	New: func() any {
		return new(TmallServiceSettleadjustmentRequestResult)
	},
}

// GetTmallServiceSettleadjustmentRequestResult() 从对象池中获取TmallServiceSettleadjustmentRequestResult
func GetTmallServiceSettleadjustmentRequestResult() *TmallServiceSettleadjustmentRequestResult {
	return poolTmallServiceSettleadjustmentRequestResult.Get().(*TmallServiceSettleadjustmentRequestResult)
}

// ReleaseTmallServiceSettleadjustmentRequestResult 释放TmallServiceSettleadjustmentRequestResult
func ReleaseTmallServiceSettleadjustmentRequestResult(v *TmallServiceSettleadjustmentRequestResult) {
	v.DataModule = nil
	v.ErrorMessage = nil
	v.Success = false
	poolTmallServiceSettleadjustmentRequestResult.Put(v)
}

package tmallservice

import (
	"sync"
)

// TmallServiceSettleadjustmentCancelResult 结构体
type TmallServiceSettleadjustmentCancelResult struct {
	// errorMessage
	ErrorMessage *ErrorMessage `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServiceSettleadjustmentCancelResult = sync.Pool{
	New: func() any {
		return new(TmallServiceSettleadjustmentCancelResult)
	},
}

// GetTmallServiceSettleadjustmentCancelResult() 从对象池中获取TmallServiceSettleadjustmentCancelResult
func GetTmallServiceSettleadjustmentCancelResult() *TmallServiceSettleadjustmentCancelResult {
	return poolTmallServiceSettleadjustmentCancelResult.Get().(*TmallServiceSettleadjustmentCancelResult)
}

// ReleaseTmallServiceSettleadjustmentCancelResult 释放TmallServiceSettleadjustmentCancelResult
func ReleaseTmallServiceSettleadjustmentCancelResult(v *TmallServiceSettleadjustmentCancelResult) {
	v.ErrorMessage = nil
	v.Success = false
	poolTmallServiceSettleadjustmentCancelResult.Put(v)
}

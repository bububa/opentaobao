package user

import (
	"sync"
)

// TmallServiceSettleadjustmentModifyResult 结构体
type TmallServiceSettleadjustmentModifyResult struct {
	// errorMessage
	ErrorMessage *ErrorMessage `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServiceSettleadjustmentModifyResult = sync.Pool{
	New: func() any {
		return new(TmallServiceSettleadjustmentModifyResult)
	},
}

// GetTmallServiceSettleadjustmentModifyResult() 从对象池中获取TmallServiceSettleadjustmentModifyResult
func GetTmallServiceSettleadjustmentModifyResult() *TmallServiceSettleadjustmentModifyResult {
	return poolTmallServiceSettleadjustmentModifyResult.Get().(*TmallServiceSettleadjustmentModifyResult)
}

// ReleaseTmallServiceSettleadjustmentModifyResult 释放TmallServiceSettleadjustmentModifyResult
func ReleaseTmallServiceSettleadjustmentModifyResult(v *TmallServiceSettleadjustmentModifyResult) {
	v.ErrorMessage = nil
	v.Success = false
	poolTmallServiceSettleadjustmentModifyResult.Put(v)
}

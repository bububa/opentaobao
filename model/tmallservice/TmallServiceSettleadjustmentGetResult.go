package tmallservice

import (
	"sync"
)

// TmallServiceSettleadjustmentGetResult 结构体
type TmallServiceSettleadjustmentGetResult struct {
	// dataModule
	DataModule *SettleAdjustmentResponse `json:"data_module,omitempty" xml:"data_module,omitempty"`
	// errorMessage
	ErrorMessage *ErrorMessage `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// success，成功true，失败false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServiceSettleadjustmentGetResult = sync.Pool{
	New: func() any {
		return new(TmallServiceSettleadjustmentGetResult)
	},
}

// GetTmallServiceSettleadjustmentGetResult() 从对象池中获取TmallServiceSettleadjustmentGetResult
func GetTmallServiceSettleadjustmentGetResult() *TmallServiceSettleadjustmentGetResult {
	return poolTmallServiceSettleadjustmentGetResult.Get().(*TmallServiceSettleadjustmentGetResult)
}

// ReleaseTmallServiceSettleadjustmentGetResult 释放TmallServiceSettleadjustmentGetResult
func ReleaseTmallServiceSettleadjustmentGetResult(v *TmallServiceSettleadjustmentGetResult) {
	v.DataModule = nil
	v.ErrorMessage = nil
	v.Success = false
	poolTmallServiceSettleadjustmentGetResult.Put(v)
}

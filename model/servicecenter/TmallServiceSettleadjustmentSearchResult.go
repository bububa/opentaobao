package servicecenter

import (
	"sync"
)

// TmallServiceSettleadjustmentSearchResult 结构体
type TmallServiceSettleadjustmentSearchResult struct {
	// dataModule
	SettleAdjustmentList []SettleAdjustmentResponse `json:"settle_adjustment_list,omitempty" xml:"settle_adjustment_list>settle_adjustment_response,omitempty"`
	// errorMessage
	ErrorMessage *ErrorMessage `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServiceSettleadjustmentSearchResult = sync.Pool{
	New: func() any {
		return new(TmallServiceSettleadjustmentSearchResult)
	},
}

// GetTmallServiceSettleadjustmentSearchResult() 从对象池中获取TmallServiceSettleadjustmentSearchResult
func GetTmallServiceSettleadjustmentSearchResult() *TmallServiceSettleadjustmentSearchResult {
	return poolTmallServiceSettleadjustmentSearchResult.Get().(*TmallServiceSettleadjustmentSearchResult)
}

// ReleaseTmallServiceSettleadjustmentSearchResult 释放TmallServiceSettleadjustmentSearchResult
func ReleaseTmallServiceSettleadjustmentSearchResult(v *TmallServiceSettleadjustmentSearchResult) {
	v.SettleAdjustmentList = v.SettleAdjustmentList[:0]
	v.ErrorMessage = nil
	v.Success = false
	poolTmallServiceSettleadjustmentSearchResult.Put(v)
}

package tmallservice

import (
	"sync"
)

// TmallServiceSettleadjustmentOperateResult 结构体
type TmallServiceSettleadjustmentOperateResult struct {
	// 错误信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误提示
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServiceSettleadjustmentOperateResult = sync.Pool{
	New: func() any {
		return new(TmallServiceSettleadjustmentOperateResult)
	},
}

// GetTmallServiceSettleadjustmentOperateResult() 从对象池中获取TmallServiceSettleadjustmentOperateResult
func GetTmallServiceSettleadjustmentOperateResult() *TmallServiceSettleadjustmentOperateResult {
	return poolTmallServiceSettleadjustmentOperateResult.Get().(*TmallServiceSettleadjustmentOperateResult)
}

// ReleaseTmallServiceSettleadjustmentOperateResult 释放TmallServiceSettleadjustmentOperateResult
func ReleaseTmallServiceSettleadjustmentOperateResult(v *TmallServiceSettleadjustmentOperateResult) {
	v.DisplayMsg = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolTmallServiceSettleadjustmentOperateResult.Put(v)
}

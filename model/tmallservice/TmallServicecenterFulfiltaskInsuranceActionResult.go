package tmallservice

import (
	"sync"
)

// TmallServicecenterFulfiltaskInsuranceActionResult 结构体
type TmallServicecenterFulfiltaskInsuranceActionResult struct {
	// 对外展示的错误信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回值对象
	Value *FulfilInsuranceTaskResponse `json:"value,omitempty" xml:"value,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServicecenterFulfiltaskInsuranceActionResult = sync.Pool{
	New: func() any {
		return new(TmallServicecenterFulfiltaskInsuranceActionResult)
	},
}

// GetTmallServicecenterFulfiltaskInsuranceActionResult() 从对象池中获取TmallServicecenterFulfiltaskInsuranceActionResult
func GetTmallServicecenterFulfiltaskInsuranceActionResult() *TmallServicecenterFulfiltaskInsuranceActionResult {
	return poolTmallServicecenterFulfiltaskInsuranceActionResult.Get().(*TmallServicecenterFulfiltaskInsuranceActionResult)
}

// ReleaseTmallServicecenterFulfiltaskInsuranceActionResult 释放TmallServicecenterFulfiltaskInsuranceActionResult
func ReleaseTmallServicecenterFulfiltaskInsuranceActionResult(v *TmallServicecenterFulfiltaskInsuranceActionResult) {
	v.DisplayMsg = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Value = nil
	v.Success = false
	poolTmallServicecenterFulfiltaskInsuranceActionResult.Put(v)
}

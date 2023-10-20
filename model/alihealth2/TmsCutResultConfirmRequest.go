package alihealth2

import (
	"sync"
)

// TmsCutResultConfirmRequest 结构体
type TmsCutResultConfirmRequest struct {
	// 回传数组
	ConfirmItems []TmsCutResultConfirmItem `json:"confirm_items,omitempty" xml:"confirm_items>tms_cut_result_confirm_item,omitempty"`
	// 扩展参数
	Extend string `json:"extend,omitempty" xml:"extend,omitempty"`
}

var poolTmsCutResultConfirmRequest = sync.Pool{
	New: func() any {
		return new(TmsCutResultConfirmRequest)
	},
}

// GetTmsCutResultConfirmRequest() 从对象池中获取TmsCutResultConfirmRequest
func GetTmsCutResultConfirmRequest() *TmsCutResultConfirmRequest {
	return poolTmsCutResultConfirmRequest.Get().(*TmsCutResultConfirmRequest)
}

// ReleaseTmsCutResultConfirmRequest 释放TmsCutResultConfirmRequest
func ReleaseTmsCutResultConfirmRequest(v *TmsCutResultConfirmRequest) {
	v.ConfirmItems = v.ConfirmItems[:0]
	v.Extend = ""
	poolTmsCutResultConfirmRequest.Put(v)
}

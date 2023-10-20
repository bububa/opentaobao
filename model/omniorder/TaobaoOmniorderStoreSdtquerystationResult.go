package omniorder

import (
	"sync"
)

// TaobaoOmniorderStoreSdtquerystationResult 结构体
type TaobaoOmniorderStoreSdtquerystationResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// data
	Data *SdtQueryPackageResponse `json:"data,omitempty" xml:"data,omitempty"`
}

var poolTaobaoOmniorderStoreSdtquerystationResult = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderStoreSdtquerystationResult)
	},
}

// GetTaobaoOmniorderStoreSdtquerystationResult() 从对象池中获取TaobaoOmniorderStoreSdtquerystationResult
func GetTaobaoOmniorderStoreSdtquerystationResult() *TaobaoOmniorderStoreSdtquerystationResult {
	return poolTaobaoOmniorderStoreSdtquerystationResult.Get().(*TaobaoOmniorderStoreSdtquerystationResult)
}

// ReleaseTaobaoOmniorderStoreSdtquerystationResult 释放TaobaoOmniorderStoreSdtquerystationResult
func ReleaseTaobaoOmniorderStoreSdtquerystationResult(v *TaobaoOmniorderStoreSdtquerystationResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Data = nil
	poolTaobaoOmniorderStoreSdtquerystationResult.Put(v)
}

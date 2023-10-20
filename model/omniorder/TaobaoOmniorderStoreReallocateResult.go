package omniorder

import (
	"sync"
)

// TaobaoOmniorderStoreReallocateResult 结构体
type TaobaoOmniorderStoreReallocateResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// data
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}

var poolTaobaoOmniorderStoreReallocateResult = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderStoreReallocateResult)
	},
}

// GetTaobaoOmniorderStoreReallocateResult() 从对象池中获取TaobaoOmniorderStoreReallocateResult
func GetTaobaoOmniorderStoreReallocateResult() *TaobaoOmniorderStoreReallocateResult {
	return poolTaobaoOmniorderStoreReallocateResult.Get().(*TaobaoOmniorderStoreReallocateResult)
}

// ReleaseTaobaoOmniorderStoreReallocateResult 释放TaobaoOmniorderStoreReallocateResult
func ReleaseTaobaoOmniorderStoreReallocateResult(v *TaobaoOmniorderStoreReallocateResult) {
	v.Message = ""
	v.Data = ""
	v.Code = ""
	poolTaobaoOmniorderStoreReallocateResult.Put(v)
}

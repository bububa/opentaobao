package omniorder

import (
	"sync"
)

// TaobaoOmniorderStoreCollectconfigUpdateResult 结构体
type TaobaoOmniorderStoreCollectconfigUpdateResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// data
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}

var poolTaobaoOmniorderStoreCollectconfigUpdateResult = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderStoreCollectconfigUpdateResult)
	},
}

// GetTaobaoOmniorderStoreCollectconfigUpdateResult() 从对象池中获取TaobaoOmniorderStoreCollectconfigUpdateResult
func GetTaobaoOmniorderStoreCollectconfigUpdateResult() *TaobaoOmniorderStoreCollectconfigUpdateResult {
	return poolTaobaoOmniorderStoreCollectconfigUpdateResult.Get().(*TaobaoOmniorderStoreCollectconfigUpdateResult)
}

// ReleaseTaobaoOmniorderStoreCollectconfigUpdateResult 释放TaobaoOmniorderStoreCollectconfigUpdateResult
func ReleaseTaobaoOmniorderStoreCollectconfigUpdateResult(v *TaobaoOmniorderStoreCollectconfigUpdateResult) {
	v.Message = ""
	v.Data = ""
	v.Code = ""
	poolTaobaoOmniorderStoreCollectconfigUpdateResult.Put(v)
}

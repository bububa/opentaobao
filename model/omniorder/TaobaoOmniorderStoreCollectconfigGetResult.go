package omniorder

import (
	"sync"
)

// TaobaoOmniorderStoreCollectconfigGetResult 结构体
type TaobaoOmniorderStoreCollectconfigGetResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data *StoreCollectConfig `json:"data,omitempty" xml:"data,omitempty"`
}

var poolTaobaoOmniorderStoreCollectconfigGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderStoreCollectconfigGetResult)
	},
}

// GetTaobaoOmniorderStoreCollectconfigGetResult() 从对象池中获取TaobaoOmniorderStoreCollectconfigGetResult
func GetTaobaoOmniorderStoreCollectconfigGetResult() *TaobaoOmniorderStoreCollectconfigGetResult {
	return poolTaobaoOmniorderStoreCollectconfigGetResult.Get().(*TaobaoOmniorderStoreCollectconfigGetResult)
}

// ReleaseTaobaoOmniorderStoreCollectconfigGetResult 释放TaobaoOmniorderStoreCollectconfigGetResult
func ReleaseTaobaoOmniorderStoreCollectconfigGetResult(v *TaobaoOmniorderStoreCollectconfigGetResult) {
	v.Message = ""
	v.Code = ""
	v.Data = nil
	poolTaobaoOmniorderStoreCollectconfigGetResult.Put(v)
}

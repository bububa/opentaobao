package omniorder

import (
	"sync"
)

// TaobaoOmniorderStoreDeliverconfigGetResult 结构体
type TaobaoOmniorderStoreDeliverconfigGetResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data *StoreDeliverConfig `json:"data,omitempty" xml:"data,omitempty"`
}

var poolTaobaoOmniorderStoreDeliverconfigGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderStoreDeliverconfigGetResult)
	},
}

// GetTaobaoOmniorderStoreDeliverconfigGetResult() 从对象池中获取TaobaoOmniorderStoreDeliverconfigGetResult
func GetTaobaoOmniorderStoreDeliverconfigGetResult() *TaobaoOmniorderStoreDeliverconfigGetResult {
	return poolTaobaoOmniorderStoreDeliverconfigGetResult.Get().(*TaobaoOmniorderStoreDeliverconfigGetResult)
}

// ReleaseTaobaoOmniorderStoreDeliverconfigGetResult 释放TaobaoOmniorderStoreDeliverconfigGetResult
func ReleaseTaobaoOmniorderStoreDeliverconfigGetResult(v *TaobaoOmniorderStoreDeliverconfigGetResult) {
	v.Message = ""
	v.Code = ""
	v.Data = nil
	poolTaobaoOmniorderStoreDeliverconfigGetResult.Put(v)
}

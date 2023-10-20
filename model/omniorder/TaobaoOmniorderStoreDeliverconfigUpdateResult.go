package omniorder

import (
	"sync"
)

// TaobaoOmniorderStoreDeliverconfigUpdateResult 结构体
type TaobaoOmniorderStoreDeliverconfigUpdateResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// data
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}

var poolTaobaoOmniorderStoreDeliverconfigUpdateResult = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderStoreDeliverconfigUpdateResult)
	},
}

// GetTaobaoOmniorderStoreDeliverconfigUpdateResult() 从对象池中获取TaobaoOmniorderStoreDeliverconfigUpdateResult
func GetTaobaoOmniorderStoreDeliverconfigUpdateResult() *TaobaoOmniorderStoreDeliverconfigUpdateResult {
	return poolTaobaoOmniorderStoreDeliverconfigUpdateResult.Get().(*TaobaoOmniorderStoreDeliverconfigUpdateResult)
}

// ReleaseTaobaoOmniorderStoreDeliverconfigUpdateResult 释放TaobaoOmniorderStoreDeliverconfigUpdateResult
func ReleaseTaobaoOmniorderStoreDeliverconfigUpdateResult(v *TaobaoOmniorderStoreDeliverconfigUpdateResult) {
	v.Message = ""
	v.Data = ""
	v.Code = ""
	poolTaobaoOmniorderStoreDeliverconfigUpdateResult.Put(v)
}

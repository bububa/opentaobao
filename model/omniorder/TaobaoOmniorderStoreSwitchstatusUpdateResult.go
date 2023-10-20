package omniorder

import (
	"sync"
)

// TaobaoOmniorderStoreSwitchstatusUpdateResult 结构体
type TaobaoOmniorderStoreSwitchstatusUpdateResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// data
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}

var poolTaobaoOmniorderStoreSwitchstatusUpdateResult = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderStoreSwitchstatusUpdateResult)
	},
}

// GetTaobaoOmniorderStoreSwitchstatusUpdateResult() 从对象池中获取TaobaoOmniorderStoreSwitchstatusUpdateResult
func GetTaobaoOmniorderStoreSwitchstatusUpdateResult() *TaobaoOmniorderStoreSwitchstatusUpdateResult {
	return poolTaobaoOmniorderStoreSwitchstatusUpdateResult.Get().(*TaobaoOmniorderStoreSwitchstatusUpdateResult)
}

// ReleaseTaobaoOmniorderStoreSwitchstatusUpdateResult 释放TaobaoOmniorderStoreSwitchstatusUpdateResult
func ReleaseTaobaoOmniorderStoreSwitchstatusUpdateResult(v *TaobaoOmniorderStoreSwitchstatusUpdateResult) {
	v.Message = ""
	v.Data = ""
	v.Code = ""
	poolTaobaoOmniorderStoreSwitchstatusUpdateResult.Put(v)
}

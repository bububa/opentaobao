package omniorder

import (
	"sync"
)

// TaobaoOmniorderStoreSwitchstatusGetResult 结构体
type TaobaoOmniorderStoreSwitchstatusGetResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回素材id
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}

var poolTaobaoOmniorderStoreSwitchstatusGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderStoreSwitchstatusGetResult)
	},
}

// GetTaobaoOmniorderStoreSwitchstatusGetResult() 从对象池中获取TaobaoOmniorderStoreSwitchstatusGetResult
func GetTaobaoOmniorderStoreSwitchstatusGetResult() *TaobaoOmniorderStoreSwitchstatusGetResult {
	return poolTaobaoOmniorderStoreSwitchstatusGetResult.Get().(*TaobaoOmniorderStoreSwitchstatusGetResult)
}

// ReleaseTaobaoOmniorderStoreSwitchstatusGetResult 释放TaobaoOmniorderStoreSwitchstatusGetResult
func ReleaseTaobaoOmniorderStoreSwitchstatusGetResult(v *TaobaoOmniorderStoreSwitchstatusGetResult) {
	v.Message = ""
	v.Data = ""
	v.Code = ""
	poolTaobaoOmniorderStoreSwitchstatusGetResult.Put(v)
}

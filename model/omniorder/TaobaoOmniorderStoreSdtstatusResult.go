package omniorder

import (
	"sync"
)

// TaobaoOmniorderStoreSdtstatusResult 结构体
type TaobaoOmniorderStoreSdtstatusResult struct {
	// 异常信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 异常码 0  正常，否则异常
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// data
	Data *SdtStatusResponse `json:"data,omitempty" xml:"data,omitempty"`
}

var poolTaobaoOmniorderStoreSdtstatusResult = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderStoreSdtstatusResult)
	},
}

// GetTaobaoOmniorderStoreSdtstatusResult() 从对象池中获取TaobaoOmniorderStoreSdtstatusResult
func GetTaobaoOmniorderStoreSdtstatusResult() *TaobaoOmniorderStoreSdtstatusResult {
	return poolTaobaoOmniorderStoreSdtstatusResult.Get().(*TaobaoOmniorderStoreSdtstatusResult)
}

// ReleaseTaobaoOmniorderStoreSdtstatusResult 释放TaobaoOmniorderStoreSdtstatusResult
func ReleaseTaobaoOmniorderStoreSdtstatusResult(v *TaobaoOmniorderStoreSdtstatusResult) {
	v.Message = ""
	v.ErrCode = ""
	v.Data = nil
	poolTaobaoOmniorderStoreSdtstatusResult.Put(v)
}

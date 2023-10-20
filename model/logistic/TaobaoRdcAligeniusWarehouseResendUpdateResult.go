package logistic

import (
	"sync"
)

// TaobaoRdcAligeniusWarehouseResendUpdateResult 结构体
type TaobaoRdcAligeniusWarehouseResendUpdateResult struct {
	// errorInfo
	FailInfo string `json:"fail_info,omitempty" xml:"fail_info,omitempty"`
	// errorCode
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// success
	SuccessFlag bool `json:"success_flag,omitempty" xml:"success_flag,omitempty"`
}

var poolTaobaoRdcAligeniusWarehouseResendUpdateResult = sync.Pool{
	New: func() any {
		return new(TaobaoRdcAligeniusWarehouseResendUpdateResult)
	},
}

// GetTaobaoRdcAligeniusWarehouseResendUpdateResult() 从对象池中获取TaobaoRdcAligeniusWarehouseResendUpdateResult
func GetTaobaoRdcAligeniusWarehouseResendUpdateResult() *TaobaoRdcAligeniusWarehouseResendUpdateResult {
	return poolTaobaoRdcAligeniusWarehouseResendUpdateResult.Get().(*TaobaoRdcAligeniusWarehouseResendUpdateResult)
}

// ReleaseTaobaoRdcAligeniusWarehouseResendUpdateResult 释放TaobaoRdcAligeniusWarehouseResendUpdateResult
func ReleaseTaobaoRdcAligeniusWarehouseResendUpdateResult(v *TaobaoRdcAligeniusWarehouseResendUpdateResult) {
	v.FailInfo = ""
	v.FailCode = ""
	v.SuccessFlag = false
	poolTaobaoRdcAligeniusWarehouseResendUpdateResult.Put(v)
}

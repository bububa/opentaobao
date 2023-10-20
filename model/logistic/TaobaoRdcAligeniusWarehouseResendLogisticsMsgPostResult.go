package logistic

import (
	"sync"
)

// TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResult 结构体
type TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResult struct {
	// 错误信息
	FailInfo string `json:"fail_info,omitempty" xml:"fail_info,omitempty"`
	// 错误code
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// success
	SuccessFlag bool `json:"success_flag,omitempty" xml:"success_flag,omitempty"`
}

var poolTaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResult = sync.Pool{
	New: func() any {
		return new(TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResult)
	},
}

// GetTaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResult() 从对象池中获取TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResult
func GetTaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResult() *TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResult {
	return poolTaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResult.Get().(*TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResult)
}

// ReleaseTaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResult 释放TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResult
func ReleaseTaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResult(v *TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResult) {
	v.FailInfo = ""
	v.FailCode = ""
	v.SuccessFlag = false
	poolTaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResult.Put(v)
}

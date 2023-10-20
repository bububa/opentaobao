package logistic

import (
	"sync"
)

// TaobaoRdcAligeniusWarehouseReverseEventUpdateResult 结构体
type TaobaoRdcAligeniusWarehouseReverseEventUpdateResult struct {
	// 错误描述
	FailInfo string `json:"fail_info,omitempty" xml:"fail_info,omitempty"`
	// 错误码
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// 是否成功
	SuccessFlag bool `json:"success_flag,omitempty" xml:"success_flag,omitempty"`
}

var poolTaobaoRdcAligeniusWarehouseReverseEventUpdateResult = sync.Pool{
	New: func() any {
		return new(TaobaoRdcAligeniusWarehouseReverseEventUpdateResult)
	},
}

// GetTaobaoRdcAligeniusWarehouseReverseEventUpdateResult() 从对象池中获取TaobaoRdcAligeniusWarehouseReverseEventUpdateResult
func GetTaobaoRdcAligeniusWarehouseReverseEventUpdateResult() *TaobaoRdcAligeniusWarehouseReverseEventUpdateResult {
	return poolTaobaoRdcAligeniusWarehouseReverseEventUpdateResult.Get().(*TaobaoRdcAligeniusWarehouseReverseEventUpdateResult)
}

// ReleaseTaobaoRdcAligeniusWarehouseReverseEventUpdateResult 释放TaobaoRdcAligeniusWarehouseReverseEventUpdateResult
func ReleaseTaobaoRdcAligeniusWarehouseReverseEventUpdateResult(v *TaobaoRdcAligeniusWarehouseReverseEventUpdateResult) {
	v.FailInfo = ""
	v.FailCode = ""
	v.SuccessFlag = false
	poolTaobaoRdcAligeniusWarehouseReverseEventUpdateResult.Put(v)
}

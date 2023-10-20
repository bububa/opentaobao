package trade

import (
	"sync"
)

// TaobaoLifeReservationItemOrderConfirmResult 结构体
type TaobaoLifeReservationItemOrderConfirmResult struct {
	// 内部trace 用于排查问题
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误
	Error *TribeError `json:"error,omitempty" xml:"error,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoLifeReservationItemOrderConfirmResult = sync.Pool{
	New: func() any {
		return new(TaobaoLifeReservationItemOrderConfirmResult)
	},
}

// GetTaobaoLifeReservationItemOrderConfirmResult() 从对象池中获取TaobaoLifeReservationItemOrderConfirmResult
func GetTaobaoLifeReservationItemOrderConfirmResult() *TaobaoLifeReservationItemOrderConfirmResult {
	return poolTaobaoLifeReservationItemOrderConfirmResult.Get().(*TaobaoLifeReservationItemOrderConfirmResult)
}

// ReleaseTaobaoLifeReservationItemOrderConfirmResult 释放TaobaoLifeReservationItemOrderConfirmResult
func ReleaseTaobaoLifeReservationItemOrderConfirmResult(v *TaobaoLifeReservationItemOrderConfirmResult) {
	v.TraceId = ""
	v.Error = nil
	v.Success = false
	poolTaobaoLifeReservationItemOrderConfirmResult.Put(v)
}

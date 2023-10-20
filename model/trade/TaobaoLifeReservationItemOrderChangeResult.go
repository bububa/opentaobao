package trade

import (
	"sync"
)

// TaobaoLifeReservationItemOrderChangeResult 结构体
type TaobaoLifeReservationItemOrderChangeResult struct {
	// 内部trace 用于排查问题
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误
	Error *TribeError `json:"error,omitempty" xml:"error,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoLifeReservationItemOrderChangeResult = sync.Pool{
	New: func() any {
		return new(TaobaoLifeReservationItemOrderChangeResult)
	},
}

// GetTaobaoLifeReservationItemOrderChangeResult() 从对象池中获取TaobaoLifeReservationItemOrderChangeResult
func GetTaobaoLifeReservationItemOrderChangeResult() *TaobaoLifeReservationItemOrderChangeResult {
	return poolTaobaoLifeReservationItemOrderChangeResult.Get().(*TaobaoLifeReservationItemOrderChangeResult)
}

// ReleaseTaobaoLifeReservationItemOrderChangeResult 释放TaobaoLifeReservationItemOrderChangeResult
func ReleaseTaobaoLifeReservationItemOrderChangeResult(v *TaobaoLifeReservationItemOrderChangeResult) {
	v.TraceId = ""
	v.Error = nil
	v.Success = false
	poolTaobaoLifeReservationItemOrderChangeResult.Put(v)
}

package trade

import (
	"sync"
)

// TaobaoLifeReservationTradeConsumeNoticeResult 结构体
type TaobaoLifeReservationTradeConsumeNoticeResult struct {
	// traceId
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误
	Error *TribeError `json:"error,omitempty" xml:"error,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoLifeReservationTradeConsumeNoticeResult = sync.Pool{
	New: func() any {
		return new(TaobaoLifeReservationTradeConsumeNoticeResult)
	},
}

// GetTaobaoLifeReservationTradeConsumeNoticeResult() 从对象池中获取TaobaoLifeReservationTradeConsumeNoticeResult
func GetTaobaoLifeReservationTradeConsumeNoticeResult() *TaobaoLifeReservationTradeConsumeNoticeResult {
	return poolTaobaoLifeReservationTradeConsumeNoticeResult.Get().(*TaobaoLifeReservationTradeConsumeNoticeResult)
}

// ReleaseTaobaoLifeReservationTradeConsumeNoticeResult 释放TaobaoLifeReservationTradeConsumeNoticeResult
func ReleaseTaobaoLifeReservationTradeConsumeNoticeResult(v *TaobaoLifeReservationTradeConsumeNoticeResult) {
	v.TraceId = ""
	v.Error = nil
	v.Success = false
	poolTaobaoLifeReservationTradeConsumeNoticeResult.Put(v)
}

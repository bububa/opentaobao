package trade

// TaobaoLifeReservationTradeConsumeNoticeResult 结构体
type TaobaoLifeReservationTradeConsumeNoticeResult struct {
	// traceId
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误
	Error *TribeError `json:"error,omitempty" xml:"error,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

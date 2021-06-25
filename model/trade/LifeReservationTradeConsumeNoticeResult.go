package trade

// LifeReservationTradeConsumeNoticeResult 
type LifeReservationTradeConsumeNoticeResult struct {

    // traceId
    TraceId   string `json:"trace_id,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 错误
    Error   *TribeError `json:"error,omitempty"`

}

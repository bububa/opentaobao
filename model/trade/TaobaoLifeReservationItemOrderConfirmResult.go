package trade

// TaobaoLifeReservationItemOrderConfirmResult 
type TaobaoLifeReservationItemOrderConfirmResult struct {

    // 内部trace 用于排查问题
    TraceId   string `json:"trace_id,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 错误
    Error   *TribeError `json:"error,omitempty"`

}

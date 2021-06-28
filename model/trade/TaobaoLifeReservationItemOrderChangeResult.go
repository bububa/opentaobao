package trade

// TaobaoLifeReservationItemOrderChangeResult 
type TaobaoLifeReservationItemOrderChangeResult struct {

    // 内部trace 用于排查问题
    
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 错误
    
    Error   *TribeError `json:"error,omitempty" xml:"error,omitempty"`
    

}

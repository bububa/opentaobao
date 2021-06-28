package trade

// TaobaoLifeReservationItemOrderChangeResult 
/* model for simplify = false
type TaobaoLifeReservationItemOrderChangeResult struct {

    // 内部trace 用于排查问题
    
    TraceId   string `json:"trace_id,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 错误
    
    Error  *struct {
        TribeError  *TribeError `json:"tribe_error,omitempty"`
    } `json:"error,omitempty"`
    

}
*/

// TaobaoLifeReservationItemOrderChangeResult 
type TaobaoLifeReservationItemOrderChangeResult struct {

    // 内部trace 用于排查问题
    TraceId   string `json:"trace_id,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 错误
    Error   *TribeError `json:"error,omitempty"`

}

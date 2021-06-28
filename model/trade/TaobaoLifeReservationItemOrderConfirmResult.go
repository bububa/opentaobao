package trade

// TaobaoLifeReservationItemOrderConfirmResult 
/* model for simplify = false
type TaobaoLifeReservationItemOrderConfirmResult struct {

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

// TaobaoLifeReservationItemOrderConfirmResult 
type TaobaoLifeReservationItemOrderConfirmResult struct {

    // 内部trace 用于排查问题
    TraceId   string `json:"trace_id,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 错误
    Error   *TribeError `json:"error,omitempty"`

}

package train

// ChangeHoldSeatParam 
/* model for simplify = false
type ChangeHoldSeatParam struct {

    // 改签占座失败错误码
    
    ErrorCode   int64 `json:"error_code,omitempty"`
    

    // 改签单信息
    
    ApplyOrderBaseDo  *struct {
        ApplyOrderInfoDo  *ApplyOrderInfoDo `json:"apply_order_info_do,omitempty"`
    } `json:"apply_order_base_do,omitempty"`
    

    // 占座结果 true:占座成功 false:占座失败
    
    HoldSeatStatus   bool `json:"hold_seat_status,omitempty"`
    

}
*/

// ChangeHoldSeatParam 
type ChangeHoldSeatParam struct {

    // 改签占座失败错误码
    ErrorCode   int64 `json:"error_code,omitempty"`

    // 改签单信息
    ApplyOrderBaseDo   *ApplyOrderInfoDo `json:"apply_order_base_do,omitempty"`

    // 占座结果 true:占座成功 false:占座失败
    HoldSeatStatus   bool `json:"hold_seat_status,omitempty"`

}

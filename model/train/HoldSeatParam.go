package train

// HoldSeatParam 
/* model for simplify = false
type HoldSeatParam struct {

    // 订单类型
    
    OrderType   int64 `json:"order_type,omitempty"`
    

    // 和出票失败错误码回填相同，666666代表占座失败转后占座
    
    ErrorCode   int64 `json:"error_code,omitempty"`
    

    // 订单信息
    
    OrderBase  *struct {
        OrderInfoDto  *OrderInfoDto `json:"order_info_dto,omitempty"`
    } `json:"order_base,omitempty"`
    

    // 是否占座成功
    
    HoldSeatStatus   bool `json:"hold_seat_status,omitempty"`
    

    // 代理商id
    
    AgentId   int64 `json:"agent_id,omitempty"`
    

}
*/

// HoldSeatParam 
type HoldSeatParam struct {

    // 订单类型
    OrderType   int64 `json:"order_type,omitempty"`

    // 和出票失败错误码回填相同，666666代表占座失败转后占座
    ErrorCode   int64 `json:"error_code,omitempty"`

    // 订单信息
    OrderBase   *OrderInfoDto `json:"order_base,omitempty"`

    // 是否占座成功
    HoldSeatStatus   bool `json:"hold_seat_status,omitempty"`

    // 代理商id
    AgentId   int64 `json:"agent_id,omitempty"`

}

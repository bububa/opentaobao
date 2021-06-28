package jst

// RefundTrace 
/* model for simplify = false
type RefundTrace struct {

    // 交易tid
    
    Tid   int64 `json:"tid,omitempty"`
    

    // 回流的订单状态
    
    Status   string `json:"status,omitempty"`
    

    // 动作发生的时间
    
    ActionTime   string `json:"action_time,omitempty"`
    

    // 备注字段
    
    Remark   string `json:"remark,omitempty"`
    

    // 应用标题
    
    AppTitle   string `json:"app_title,omitempty"`
    

    // 退款编号
    
    RefundId   int64 `json:"refund_id,omitempty"`
    

}
*/

// RefundTrace 
type RefundTrace struct {

    // 交易tid
    Tid   int64 `json:"tid,omitempty"`

    // 回流的订单状态
    Status   string `json:"status,omitempty"`

    // 动作发生的时间
    ActionTime   string `json:"action_time,omitempty"`

    // 备注字段
    Remark   string `json:"remark,omitempty"`

    // 应用标题
    AppTitle   string `json:"app_title,omitempty"`

    // 退款编号
    RefundId   int64 `json:"refund_id,omitempty"`

}

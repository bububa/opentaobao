package util

// RefundCheckDto 
/* model for simplify = false
type RefundCheckDto struct {

    // 退款单ID
    
    RefundId   int64 `json:"refund_id,omitempty"`
    

    // 主订单ID
    
    Tid   int64 `json:"tid,omitempty"`
    

    // 子订单ID
    
    Oid   int64 `json:"oid,omitempty"`
    

    // 审核状态 恒为 SUCCESS
    
    Status   string `json:"status,omitempty"`
    

    // 审核描述
    
    Msg   string `json:"msg,omitempty"`
    

    // 审核操作时间
    
    OperateTime   string `json:"operate_time,omitempty"`
    

    // 退款金额
    
    RefundFee   int64 `json:"refund_fee,omitempty"`
    

}
*/

// RefundCheckDto 
type RefundCheckDto struct {

    // 退款单ID
    RefundId   int64 `json:"refund_id,omitempty"`

    // 主订单ID
    Tid   int64 `json:"tid,omitempty"`

    // 子订单ID
    Oid   int64 `json:"oid,omitempty"`

    // 审核状态 恒为 SUCCESS
    Status   string `json:"status,omitempty"`

    // 审核描述
    Msg   string `json:"msg,omitempty"`

    // 审核操作时间
    OperateTime   string `json:"operate_time,omitempty"`

    // 退款金额
    RefundFee   int64 `json:"refund_fee,omitempty"`

}

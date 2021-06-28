package bus

// TvmRefundApply 
/* model for simplify = false
type TvmRefundApply struct {

    // 申请单id
    
    ApplyId   int64 `json:"apply_id,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 成功时间
    
    GmtRefundSuccTime   string `json:"gmt_refund_succ_time,omitempty"`
    

    // 退款金额（分）
    
    RefundAmount   int64 `json:"refund_amount,omitempty"`
    

    // 退款状态 10(处理中) 20(已拒绝) 30(已同意) 40(已退款) 50(已受理)
    
    RefundStatus   int64 `json:"refund_status,omitempty"`
    

    // 退款批次号
    
    OutTradeNo   string `json:"out_trade_no,omitempty"`
    

}
*/

// TvmRefundApply 
type TvmRefundApply struct {

    // 申请单id
    ApplyId   int64 `json:"apply_id,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 成功时间
    GmtRefundSuccTime   string `json:"gmt_refund_succ_time,omitempty"`

    // 退款金额（分）
    RefundAmount   int64 `json:"refund_amount,omitempty"`

    // 退款状态 10(处理中) 20(已拒绝) 30(已同意) 40(已退款) 50(已受理)
    RefundStatus   int64 `json:"refund_status,omitempty"`

    // 退款批次号
    OutTradeNo   string `json:"out_trade_no,omitempty"`

}

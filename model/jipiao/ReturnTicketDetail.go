package jipiao

// ReturnTicketDetail 
type ReturnTicketDetail struct {
    // 申请单ID
    ApplyId   int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
    // 退票原因
    ApplyReasonType   int64 `json:"apply_reason_type,omitempty" xml:"apply_reason_type,omitempty"`
    // 退票申请时间
    ApplyTime   string `json:"apply_time,omitempty" xml:"apply_time,omitempty"`
    // 退票成功时间
    FirstProcessTime   string `json:"first_process_time,omitempty" xml:"first_process_time,omitempty"`
    // 订单号
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    // 退款成功时间
    PaySuccessTime   string `json:"pay_success_time,omitempty" xml:"pay_success_time,omitempty"`
    // 申退原因
    Reason   string `json:"reason,omitempty" xml:"reason,omitempty"`
    // 退款手续费（分）
    RefundFee   int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
    // 退款金额（退给买家的钱）（分）
    RefundMoney   int64 `json:"refund_money,omitempty" xml:"refund_money,omitempty"`
    // 人的费用信息
    ReturnApplyPassenge   []ReturnApplyPassenge `json:"return_apply_passenge,omitempty" xml:"return_apply_passenge>return_apply_passenge,omitempty"`
    // 申请单状态（1初始 2接受 3确认 4失败 5买家处理 6卖家处理 7等待小二回填 8退款成功）
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // creditMoney
    CreditMoney   int64 `json:"credit_money,omitempty" xml:"credit_money,omitempty"`
}

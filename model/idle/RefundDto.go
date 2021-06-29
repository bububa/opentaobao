package idle

// RefundDTO 
type RefundDTO struct {
    // 申请退款金额
    ApplyRefundFee   int64 `json:"apply_refund_fee,omitempty" xml:"apply_refund_fee,omitempty"`
    // 订单Id
    BizOrderId   int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
    // 退款单ID
    DisputeId   int64 `json:"dispute_id,omitempty" xml:"dispute_id,omitempty"`
    // 实际退款金额
    RealRefundFee   int64 `json:"real_refund_fee,omitempty" xml:"real_refund_fee,omitempty"`
    // 退款状态 与逆向DisputeStatusEnum 一致，1-9
    RefundStatus   string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
    // 退款时间
    RefundTime   string `json:"refund_time,omitempty" xml:"refund_time,omitempty"`
    // 申请退款类型: REFUND_ONLY_NO_SHIP(未发货仅退款)，REFUND_ONLY_NO_RECIVE（已发货未收到货仅退款），REFUND_ONLY_WHIT_GOODS（收到货仅退款(退货退款)）
    RefundType   string `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
    // 退货状态 GoodsStatusEnum 一致
    ReturnGoodsStatus   string `json:"return_goods_status,omitempty" xml:"return_goods_status,omitempty"`
    // 到达下一个节点的超时时间点
    Timeout   int64 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

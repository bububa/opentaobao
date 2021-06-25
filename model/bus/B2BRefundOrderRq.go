package bus

// B2BRefundOrderRq 
type B2BRefundOrderRq struct {

    // 退票原因
    RefundReason   string `json:"refund_reason,omitempty"`

    // 卖家淘宝ID
    SellerAgentId   int64 `json:"seller_agent_id,omitempty"`

    // 飞猪订单号
    AliTripOrderId   string `json:"ali_trip_order_id,omitempty"`

    // 飞猪子订单号
    SubOrderIds   []Number `json:"sub_order_ids,omitempty"`

}

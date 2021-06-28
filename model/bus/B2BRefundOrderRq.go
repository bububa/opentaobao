package bus

// B2BRefundOrderRq 
/* model for simplify = false
type B2BRefundOrderRq struct {

    // 退票原因
    
    RefundReason   string `json:"refund_reason,omitempty"`
    

    // 卖家淘宝ID
    
    SellerAgentId   int64 `json:"seller_agent_id,omitempty"`
    

    // 飞猪订单号
    
    AliTripOrderId   string `json:"ali_trip_order_id,omitempty"`
    

    // 飞猪子订单号
    
    SubOrderIds  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"sub_order_ids,omitempty"`
    

}
*/

// B2BRefundOrderRq 
type B2BRefundOrderRq struct {

    // 退票原因
    RefundReason   string `json:"refund_reason,omitempty"`

    // 卖家淘宝ID
    SellerAgentId   int64 `json:"seller_agent_id,omitempty"`

    // 飞猪订单号
    AliTripOrderId   string `json:"ali_trip_order_id,omitempty"`

    // 飞猪子订单号
    SubOrderIds   []int64 `json:"sub_order_ids,omitempty"`

}

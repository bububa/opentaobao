package wdk

// OrderSyncRefundChannel 
/* model for simplify = false
type OrderSyncRefundChannel struct {

    // 退款金额
    
    RefundAmount   int64 `json:"refund_amount,omitempty"`
    

    // 退款渠道
    
    RefundChannel   int64 `json:"refund_channel,omitempty"`
    

    // 退款单id
    
    RefundOrderId   int64 `json:"refund_order_id,omitempty"`
    

}
*/

// OrderSyncRefundChannel 
type OrderSyncRefundChannel struct {

    // 退款金额
    RefundAmount   int64 `json:"refund_amount,omitempty"`

    // 退款渠道
    RefundChannel   int64 `json:"refund_channel,omitempty"`

    // 退款单id
    RefundOrderId   int64 `json:"refund_order_id,omitempty"`

}

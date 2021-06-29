package maitix

// MoaRefundAuditParam 
type MoaRefundAuditParam struct {

    // 大麦订单号
    
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    

    // 退票原因：0=出错票，1=客人退票，2=其它原因
    
    RefundType   int64 `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
    

}

package idle

// RecycleRefundTopRequest 
type RecycleRefundTopRequest struct {
    // 申请仅退款
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 订单号
    BizOrderId   int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}

package idle

// OrderDummySendRequest 
type OrderDummySendRequest struct {
    // 需要虚拟发货的订单号
    BizOrderId   int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}

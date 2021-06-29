package logistic

// ShippingOrderEvent 
type ShippingOrderEvent struct {
    // 运单状态,0:运单创建,10:配送商接单,20:骑手接单,80:骑手到店,30:骑手取餐,40:已完成,90:配送失败
    ShippingState   int64 `json:"shipping_state,omitempty" xml:"shipping_state,omitempty"`
    // 状态变更时间
    OccurredAt   int64 `json:"occurred_at,omitempty" xml:"occurred_at,omitempty"`
}

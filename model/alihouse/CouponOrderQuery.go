package alihouse

// CouponOrderQuery 结构体
type CouponOrderQuery struct {
	// 履约单id
	TradeOrderId int64 `json:"trade_order_id,omitempty" xml:"trade_order_id,omitempty"`
}

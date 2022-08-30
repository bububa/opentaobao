package tmallcar

// SingleOrderDetailQuery 结构体
type SingleOrderDetailQuery struct {
	// 买家昵称
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

package user

// BuyerOrder 结构体
type BuyerOrder struct {
	// 买家nick
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 订单id
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
}

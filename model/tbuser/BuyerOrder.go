package tbuser

import (
	"sync"
)

// BuyerOrder 结构体
type BuyerOrder struct {
	// 买家nick
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 订单id
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
}

var poolBuyerOrder = sync.Pool{
	New: func() any {
		return new(BuyerOrder)
	},
}

// GetBuyerOrder() 从对象池中获取BuyerOrder
func GetBuyerOrder() *BuyerOrder {
	return poolBuyerOrder.Get().(*BuyerOrder)
}

// ReleaseBuyerOrder 释放BuyerOrder
func ReleaseBuyerOrder(v *BuyerOrder) {
	v.BuyerNick = ""
	v.Tid = 0
	poolBuyerOrder.Put(v)
}

package ascp

import (
	"sync"
)

// BindingConsignOrderGiftResult 结构体
type BindingConsignOrderGiftResult struct {
	// 绑赠列表
	GiftOrders []GiftOrder `json:"gift_orders,omitempty" xml:"gift_orders>gift_order,omitempty"`
	// 交易主单号
	TradeId string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// 翱象发货单号
	ConsignOrderCode string `json:"consign_order_code,omitempty" xml:"consign_order_code,omitempty"`
}

var poolBindingConsignOrderGiftResult = sync.Pool{
	New: func() any {
		return new(BindingConsignOrderGiftResult)
	},
}

// GetBindingConsignOrderGiftResult() 从对象池中获取BindingConsignOrderGiftResult
func GetBindingConsignOrderGiftResult() *BindingConsignOrderGiftResult {
	return poolBindingConsignOrderGiftResult.Get().(*BindingConsignOrderGiftResult)
}

// ReleaseBindingConsignOrderGiftResult 释放BindingConsignOrderGiftResult
func ReleaseBindingConsignOrderGiftResult(v *BindingConsignOrderGiftResult) {
	v.GiftOrders = v.GiftOrders[:0]
	v.TradeId = ""
	v.ConsignOrderCode = ""
	poolBindingConsignOrderGiftResult.Put(v)
}

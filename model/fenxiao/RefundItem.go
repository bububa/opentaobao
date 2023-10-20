package fenxiao

import (
	"sync"
)

// RefundItem 结构体
type RefundItem struct {
	// 退款明细ID，针对一笔退款每一个品就映射为一个明细，每一个明细有一个全局唯一的ID
	RefundItemId int64 `json:"refund_item_id,omitempty" xml:"refund_item_id,omitempty"`
	// 分销子订单号
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 退货数量
	RefundQuantity int64 `json:"refund_quantity,omitempty" xml:"refund_quantity,omitempty"`
}

var poolRefundItem = sync.Pool{
	New: func() any {
		return new(RefundItem)
	},
}

// GetRefundItem() 从对象池中获取RefundItem
func GetRefundItem() *RefundItem {
	return poolRefundItem.Get().(*RefundItem)
}

// ReleaseRefundItem 释放RefundItem
func ReleaseRefundItem(v *RefundItem) {
	v.RefundItemId = 0
	v.SubOrderId = 0
	v.RefundQuantity = 0
	poolRefundItem.Put(v)
}

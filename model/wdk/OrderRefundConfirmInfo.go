package wdk

import (
	"sync"
)

// OrderRefundConfirmInfo 结构体
type OrderRefundConfirmInfo struct {
	// 同意退款子单
	AgreeSubOrders []SubRefundConfirm `json:"agree_sub_orders,omitempty" xml:"agree_sub_orders>sub_refund_confirm,omitempty"`
	// 经营店ID
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 渠道店ID
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 盒马主单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 外部退款批次Id
	OutRefundBatchId string `json:"out_refund_batch_id,omitempty" xml:"out_refund_batch_id,omitempty"`
}

var poolOrderRefundConfirmInfo = sync.Pool{
	New: func() any {
		return new(OrderRefundConfirmInfo)
	},
}

// GetOrderRefundConfirmInfo() 从对象池中获取OrderRefundConfirmInfo
func GetOrderRefundConfirmInfo() *OrderRefundConfirmInfo {
	return poolOrderRefundConfirmInfo.Get().(*OrderRefundConfirmInfo)
}

// ReleaseOrderRefundConfirmInfo 释放OrderRefundConfirmInfo
func ReleaseOrderRefundConfirmInfo(v *OrderRefundConfirmInfo) {
	v.AgreeSubOrders = v.AgreeSubOrders[:0]
	v.StoreId = ""
	v.ShopId = ""
	v.BizOrderId = ""
	v.OutRefundBatchId = ""
	poolOrderRefundConfirmInfo.Put(v)
}

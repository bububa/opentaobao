package trade

import (
	"sync"
)

// OrderRefundOperationReq 结构体
type OrderRefundOperationReq struct {
	// 退款退货操作的Code，由系统定义，目前支持的方式有：refundFeeOnly(仅退款)，refundFeeWithGoods(退货退款),swithGoods(换货)
	RefundOrderActionType string `json:"refund_order_action_type,omitempty" xml:"refund_order_action_type,omitempty"`
	// 订单ID
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolOrderRefundOperationReq = sync.Pool{
	New: func() any {
		return new(OrderRefundOperationReq)
	},
}

// GetOrderRefundOperationReq() 从对象池中获取OrderRefundOperationReq
func GetOrderRefundOperationReq() *OrderRefundOperationReq {
	return poolOrderRefundOperationReq.Get().(*OrderRefundOperationReq)
}

// ReleaseOrderRefundOperationReq 释放OrderRefundOperationReq
func ReleaseOrderRefundOperationReq(v *OrderRefundOperationReq) {
	v.RefundOrderActionType = ""
	v.OrderId = ""
	poolOrderRefundOperationReq.Put(v)
}

package trade

import (
	"sync"
)

// OrderRefundOperationJudgementReq 结构体
type OrderRefundOperationJudgementReq struct {
	// 退款退货操作的Code，由系统定义，目前支持的方式有：refundFeeOnly(仅退款)，refundFeeWithGoods(退货退款),swithGoods(换货)
	RefundOrderActionType string `json:"refund_order_action_type,omitempty" xml:"refund_order_action_type,omitempty"`
	// 订单ID
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolOrderRefundOperationJudgementReq = sync.Pool{
	New: func() any {
		return new(OrderRefundOperationJudgementReq)
	},
}

// GetOrderRefundOperationJudgementReq() 从对象池中获取OrderRefundOperationJudgementReq
func GetOrderRefundOperationJudgementReq() *OrderRefundOperationJudgementReq {
	return poolOrderRefundOperationJudgementReq.Get().(*OrderRefundOperationJudgementReq)
}

// ReleaseOrderRefundOperationJudgementReq 释放OrderRefundOperationJudgementReq
func ReleaseOrderRefundOperationJudgementReq(v *OrderRefundOperationJudgementReq) {
	v.RefundOrderActionType = ""
	v.OrderId = ""
	poolOrderRefundOperationJudgementReq.Put(v)
}

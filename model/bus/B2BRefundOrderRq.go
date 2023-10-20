package bus

import (
	"sync"
)

// B2BRefundOrderRq 结构体
type B2BRefundOrderRq struct {
	// 飞猪子订单号
	SubOrderIds []int64 `json:"sub_order_ids,omitempty" xml:"sub_order_ids>int64,omitempty"`
	// 退票原因
	RefundReason string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// 飞猪订单号
	AliTripOrderId string `json:"ali_trip_order_id,omitempty" xml:"ali_trip_order_id,omitempty"`
	// 卖家淘宝ID
	SellerAgentId int64 `json:"seller_agent_id,omitempty" xml:"seller_agent_id,omitempty"`
}

var poolB2BRefundOrderRq = sync.Pool{
	New: func() any {
		return new(B2BRefundOrderRq)
	},
}

// GetB2BRefundOrderRq() 从对象池中获取B2BRefundOrderRq
func GetB2BRefundOrderRq() *B2BRefundOrderRq {
	return poolB2BRefundOrderRq.Get().(*B2BRefundOrderRq)
}

// ReleaseB2BRefundOrderRq 释放B2BRefundOrderRq
func ReleaseB2BRefundOrderRq(v *B2BRefundOrderRq) {
	v.SubOrderIds = v.SubOrderIds[:0]
	v.RefundReason = ""
	v.AliTripOrderId = ""
	v.SellerAgentId = 0
	poolB2BRefundOrderRq.Put(v)
}

package ieagency

import (
	"sync"
)

// RefundOrderQueryDetailRq 结构体
type RefundOrderQueryDetailRq struct {
	// 代理商ID
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 申请单ID
	RefundOrderId int64 `json:"refund_order_id,omitempty" xml:"refund_order_id,omitempty"`
}

var poolRefundOrderQueryDetailRq = sync.Pool{
	New: func() any {
		return new(RefundOrderQueryDetailRq)
	},
}

// GetRefundOrderQueryDetailRq() 从对象池中获取RefundOrderQueryDetailRq
func GetRefundOrderQueryDetailRq() *RefundOrderQueryDetailRq {
	return poolRefundOrderQueryDetailRq.Get().(*RefundOrderQueryDetailRq)
}

// ReleaseRefundOrderQueryDetailRq 释放RefundOrderQueryDetailRq
func ReleaseRefundOrderQueryDetailRq(v *RefundOrderQueryDetailRq) {
	v.AgentId = 0
	v.RefundOrderId = 0
	poolRefundOrderQueryDetailRq.Put(v)
}

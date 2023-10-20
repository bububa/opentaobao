package bus

import (
	"sync"
)

// AgentQueryOrderRq 结构体
type AgentQueryOrderRq struct {
	// 商家订单号，多个使用英文逗号进行分隔
	AgentOrderIds string `json:"agent_order_ids,omitempty" xml:"agent_order_ids,omitempty"`
	// 订单完成日期区间，格式：yyyyMMdd
	CompleteOrderDateInterval string `json:"complete_order_date_interval,omitempty" xml:"complete_order_date_interval,omitempty"`
	// 退票完成日期区间，，格式：yyyyMMdd
	CompleteRefundDateInterval string `json:"complete_refund_date_interval,omitempty" xml:"complete_refund_date_interval,omitempty"`
	// 订单创建日期区间，格式：yyyyMMdd
	CreateOrderDateInterval string `json:"create_order_date_interval,omitempty" xml:"create_order_date_interval,omitempty"`
	// 申请退票日期区间，，格式：yyyyMMdd
	CreateRefundDateInterval string `json:"create_refund_date_interval,omitempty" xml:"create_refund_date_interval,omitempty"`
	// 飞猪订单号，多个使用英文逗号进行分隔
	OrderIds string `json:"order_ids,omitempty" xml:"order_ids,omitempty"`
	// 退票申请id,多个使用英文逗号进行分隔
	RefundApplyIds string `json:"refund_apply_ids,omitempty" xml:"refund_apply_ids,omitempty"`
	// 当前页码从1开始
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 每页数量最大为100，超过100会有超时
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolAgentQueryOrderRq = sync.Pool{
	New: func() any {
		return new(AgentQueryOrderRq)
	},
}

// GetAgentQueryOrderRq() 从对象池中获取AgentQueryOrderRq
func GetAgentQueryOrderRq() *AgentQueryOrderRq {
	return poolAgentQueryOrderRq.Get().(*AgentQueryOrderRq)
}

// ReleaseAgentQueryOrderRq 释放AgentQueryOrderRq
func ReleaseAgentQueryOrderRq(v *AgentQueryOrderRq) {
	v.AgentOrderIds = ""
	v.CompleteOrderDateInterval = ""
	v.CompleteRefundDateInterval = ""
	v.CreateOrderDateInterval = ""
	v.CreateRefundDateInterval = ""
	v.OrderIds = ""
	v.RefundApplyIds = ""
	v.PageIndex = 0
	v.PageSize = 0
	poolAgentQueryOrderRq.Put(v)
}

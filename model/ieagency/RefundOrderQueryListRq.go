package ieagency

import (
	"sync"
)

// RefundOrderQueryListRq 结构体
type RefundOrderQueryListRq struct {
	// 申请单创建开始时间
	CreateEndTime string `json:"create_end_time,omitempty" xml:"create_end_time,omitempty"`
	// 申请单创建结束时间
	CreateStartTime string `json:"create_start_time,omitempty" xml:"create_start_time,omitempty"`
	// 代理商ID
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 【必填】分页索引
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 【必填】分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 申请单状态(WAIT(1,&#34;待处理&#34;), AGREED(2, &#34;已同意&#34;),REFUSE(3, &#34;已拒绝&#34;),PROCESS(6, &#34;已受理&#34;), SUCCESS(7, &#34;已退款&#34;))
	RefundStatus int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 飞猪订单ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolRefundOrderQueryListRq = sync.Pool{
	New: func() any {
		return new(RefundOrderQueryListRq)
	},
}

// GetRefundOrderQueryListRq() 从对象池中获取RefundOrderQueryListRq
func GetRefundOrderQueryListRq() *RefundOrderQueryListRq {
	return poolRefundOrderQueryListRq.Get().(*RefundOrderQueryListRq)
}

// ReleaseRefundOrderQueryListRq 释放RefundOrderQueryListRq
func ReleaseRefundOrderQueryListRq(v *RefundOrderQueryListRq) {
	v.CreateEndTime = ""
	v.CreateStartTime = ""
	v.AgentId = 0
	v.PageIndex = 0
	v.PageSize = 0
	v.RefundStatus = 0
	v.OrderId = 0
	poolRefundOrderQueryListRq.Put(v)
}

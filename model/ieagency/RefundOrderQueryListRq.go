package ieagency

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
	// 申请单状态(WAIT(1,"待处理"), AGREED(2, "已同意"),REFUSE(3, "已拒绝"),PROCESS(6, "已受理"), SUCCESS(7, "已退款"))
	RefundStatus int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 飞猪订单ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

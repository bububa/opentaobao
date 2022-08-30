package flight

// RefundListRequestDto 结构体
type RefundListRequestDto struct {
	// 店铺id集合
	AgentIds []int64 `json:"agent_ids,omitempty" xml:"agent_ids>int64,omitempty"`
	// 退票申请起始时间
	EndApplyTime string `json:"end_apply_time,omitempty" xml:"end_apply_time,omitempty"`
	// 退票申请结束时间(提交申请结束时间，限制只能与起始时间相差一天)
	BeginApplyTime string `json:"begin_apply_time,omitempty" xml:"begin_apply_time,omitempty"`
	// 页码
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 退票单状态:1:待回填费用,2:待退款,3:退款中,4:已完成,5:已拒绝
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

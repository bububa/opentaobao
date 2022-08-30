package flight

// ModifyListRequestDto 结构体
type ModifyListRequestDto struct {
	// 店铺id
	AgentIds []int64 `json:"agent_ids,omitempty" xml:"agent_ids>int64,omitempty"`
	// 申请结束时间(提交申请结束时间，限制只能与起始时间相差一天)
	EndApplyTime string `json:"end_apply_time,omitempty" xml:"end_apply_time,omitempty"`
	// 申请开始时间
	BeginApplyTime string `json:"begin_apply_time,omitempty" xml:"begin_apply_time,omitempty"`
	// 页码
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 改签单状态:1:待回填费用或行程,2:待用户支付,3:待出票,4:已完成,5:已拒绝
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

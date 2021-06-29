package flight

// ModifyListRequestDTO 
type ModifyListRequestDTO struct {
    // 店铺id
    AgentIds   []int64 `json:"agent_ids,omitempty" xml:"agent_ids>int64,omitempty"`
    // 申请结束时间
    EndApplyTime   string `json:"end_apply_time,omitempty" xml:"end_apply_time,omitempty"`
    // 申请开始时间
    BeginApplyTime   string `json:"begin_apply_time,omitempty" xml:"begin_apply_time,omitempty"`
    // 页码
    Page   int64 `json:"page,omitempty" xml:"page,omitempty"`
    // 改签单状态
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
}

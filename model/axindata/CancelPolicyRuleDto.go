package axindata

// CancelPolicyRuleDto 结构体
type CancelPolicyRuleDto struct {
	// 取消政策明细列表
	CancelPolicyDetailList []CancelPolicyDetailDto `json:"cancel_policy_detail_list,omitempty" xml:"cancel_policy_detail_list>cancel_policy_detail_dto,omitempty"`
	// 取消政策描述
	CancelPolicyDesc string `json:"cancel_policy_desc,omitempty" xml:"cancel_policy_desc,omitempty"`
	// 取消政策类型
	CancelType int64 `json:"cancel_type,omitempty" xml:"cancel_type,omitempty"`
	// 规则生效开始时间
	StartDate int64 `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 规则生效结束时间
	EndDate int64 `json:"end_date,omitempty" xml:"end_date,omitempty"`
}

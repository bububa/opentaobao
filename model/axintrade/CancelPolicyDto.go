package axintrade

// CancelPolicyDto 结构体
type CancelPolicyDto struct {
	// 详细规则
	PolicyInfoList []CancelPolicyInfoDto `json:"policy_info_list,omitempty" xml:"policy_info_list>cancel_policy_info_dto,omitempty"`
	// 取消政策类型
	CancelPolicyType int64 `json:"cancel_policy_type,omitempty" xml:"cancel_policy_type,omitempty"`
}

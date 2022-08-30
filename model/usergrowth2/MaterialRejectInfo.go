package usergrowth2

// MaterialRejectInfo 结构体
type MaterialRejectInfo struct {
	// 违规一级分类描述
	RuleDesc string `json:"rule_desc,omitempty" xml:"rule_desc,omitempty"`
	// 违规二级分类描述
	SubRuleDesc string `json:"sub_rule_desc,omitempty" xml:"sub_rule_desc,omitempty"`
	// 驳回原因
	RejectMemo string `json:"reject_memo,omitempty" xml:"reject_memo,omitempty"`
	// 违规一级分类id
	RuleId int64 `json:"rule_id,omitempty" xml:"rule_id,omitempty"`
	// 违规二级分类id
	SubRuleId int64 `json:"sub_rule_id,omitempty" xml:"sub_rule_id,omitempty"`
}

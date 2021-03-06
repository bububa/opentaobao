package wdk

// ReasonVo 结构体
type ReasonVo struct {
	// 标签
	Tags []string `json:"tags,omitempty" xml:"tags>string,omitempty"`
	// 原因描述
	ReasonText string `json:"reason_text,omitempty" xml:"reason_text,omitempty"`
	// tip
	ReasonTip string `json:"reason_tip,omitempty" xml:"reason_tip,omitempty"`
	// 原因id
	ReasonId int64 `json:"reason_id,omitempty" xml:"reason_id,omitempty"`
}

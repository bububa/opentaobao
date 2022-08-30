package alitripmerchant

// MemberRightVo 结构体
type MemberRightVo struct {
	// 权益名称
	MemberRightDesc string `json:"member_right_desc,omitempty" xml:"member_right_desc,omitempty"`
	// 权益内容
	MemberRightContent string `json:"member_right_content,omitempty" xml:"member_right_content,omitempty"`
	// 权益码点
	IconCodePoint string `json:"icon_code_point,omitempty" xml:"icon_code_point,omitempty"`
	// 顺序
	Order int64 `json:"order,omitempty" xml:"order,omitempty"`
}

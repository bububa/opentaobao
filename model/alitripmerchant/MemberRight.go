package alitripmerchant

// MemberRight 结构体
type MemberRight struct {
	// 会员权益名称
	MemberRightDesc string `json:"member_right_desc,omitempty" xml:"member_right_desc,omitempty"`
	// 会员权益内容
	MemberRightContent string `json:"member_right_content,omitempty" xml:"member_right_content,omitempty"`
	// icon 码点
	IconCodePoint string `json:"icon_code_point,omitempty" xml:"icon_code_point,omitempty"`
	// 权益顺序
	Order string `json:"order,omitempty" xml:"order,omitempty"`
	// 权益是否置灰封锁
	IfLock bool `json:"if_lock,omitempty" xml:"if_lock,omitempty"`
	// 是否标识为 new 权益
	IfNew bool `json:"if_new,omitempty" xml:"if_new,omitempty"`
}

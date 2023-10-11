package wdk

// OrderInfoExt 结构体
type OrderInfoExt struct {
	// 祝福语
	SubscribeMessage string `json:"subscribe_message,omitempty" xml:"subscribe_message,omitempty"`
	// 订购人电话, 虚拟小号; (商家自配送场景给出)
	SubscribePhone string `json:"subscribe_phone,omitempty" xml:"subscribe_phone,omitempty"`
	// 会员卡号
	MemberCardNum string `json:"member_card_num,omitempty" xml:"member_card_num,omitempty"`
}

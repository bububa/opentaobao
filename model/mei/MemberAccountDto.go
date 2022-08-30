package mei

// MemberAccountDto 结构体
type MemberAccountDto struct {
	// mixMobile，只有有权限的才有值
	MixMobile string `json:"mix_mobile,omitempty" xml:"mix_mobile,omitempty"`
	// buyerNick，只有有权限的才有值
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 明文手机号，只有有权限的才有值
	ClearMobile string `json:"clear_mobile,omitempty" xml:"clear_mobile,omitempty"`
}

package alitripmerchant

// DerbyVoucherCardActiveVo 结构体
type DerbyVoucherCardActiveVo struct {
	// 会员卡号
	CardNumber string `json:"card_number,omitempty" xml:"card_number,omitempty"`
	// 权益卡号
	MemberVoucherCardID string `json:"member_voucher_card_i_d,omitempty" xml:"member_voucher_card_i_d,omitempty"`
}

package alitripmerchant

// DerbyVoucherInfo 结构体
type DerbyVoucherInfo struct {
	// 权益券类型
	Category string `json:"category,omitempty" xml:"category,omitempty"`
	// 权益卡卡号
	MemberVoucherCardId string `json:"member_voucher_card_id,omitempty" xml:"member_voucher_card_id,omitempty"`
	// 权益券id
	MemberVoucherId string `json:"member_voucher_id,omitempty" xml:"member_voucher_id,omitempty"`
	// 权益券使用限制描述
	LongDes string `json:"long_des,omitempty" xml:"long_des,omitempty"`
	// 权益卡身份码
	QrCodeIDImage string `json:"qr_code_i_d_image,omitempty" xml:"qr_code_i_d_image,omitempty"`
	// 折扣百分比
	DiscountOff int64 `json:"discount_off,omitempty" xml:"discount_off,omitempty"`
}

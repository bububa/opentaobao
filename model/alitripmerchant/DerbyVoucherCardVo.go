package alitripmerchant

// DerbyVoucherCardVo 结构体
type DerbyVoucherCardVo struct {
	// 1
	DerbyVoucherCountVOs []DerbyVoucherCountVo `json:"derby_voucher_count_v_os,omitempty" xml:"derby_voucher_count_v_os>derby_voucher_count_vo,omitempty"`
	// 1
	DerbyVoucherPolymerizationVOs []DerbyVoucherCountVo `json:"derby_voucher_polymerization_v_os,omitempty" xml:"derby_voucher_polymerization_v_os>derby_voucher_count_vo,omitempty"`
	// 1
	Vouchers []DerbyVoucherPolymerizationVo `json:"vouchers,omitempty" xml:"vouchers>derby_voucher_polymerization_vo,omitempty"`
	// 1
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 1
	MemberVoucherCardID string `json:"member_voucher_card_i_d,omitempty" xml:"member_voucher_card_i_d,omitempty"`
	// 1
	VoucherCardCategory string `json:"voucher_card_category,omitempty" xml:"voucher_card_category,omitempty"`
	// 1
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 1
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 1
	QrCodeIDImage string `json:"qr_code_i_d_image,omitempty" xml:"qr_code_i_d_image,omitempty"`
}

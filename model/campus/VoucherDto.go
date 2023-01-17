package campus

// VoucherDto 结构体
type VoucherDto struct {
	// 卡号
	VoucherNo string `json:"voucher_no,omitempty" xml:"voucher_no,omitempty"`
	// xxx
	VoucherTypeEnum string `json:"voucher_type_enum,omitempty" xml:"voucher_type_enum,omitempty"`
}

package campus

// VoucherDto 
type VoucherDto struct {
    // 凭证类型
    UserVoucherTypeEnum   string `json:"user_voucher_type_enum,omitempty" xml:"user_voucher_type_enum,omitempty"`
    // 卡号
    VoucherNo   string `json:"voucher_no,omitempty" xml:"voucher_no,omitempty"`
}

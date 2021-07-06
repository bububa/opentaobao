package alsc

// VoucherStatusInfo 结构体
type VoucherStatusInfo struct {
	// 券码
	VoucherId string `json:"voucher_id,omitempty" xml:"voucher_id,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

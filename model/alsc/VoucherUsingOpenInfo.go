package alsc

// VoucherUsingOpenInfo 结构体
type VoucherUsingOpenInfo struct {
	// 1
	VoucherStatusList []VoucherStatus `json:"voucher_status_list,omitempty" xml:"voucher_status_list>voucher_status,omitempty"`
}

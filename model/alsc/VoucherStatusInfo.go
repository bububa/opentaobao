package alsc

// VoucherStatusInfo 
type VoucherStatusInfo struct {
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 券码
    VoucherId   string `json:"voucher_id,omitempty" xml:"voucher_id,omitempty"`
}

package alsc

// VoucherStatus 
type VoucherStatus struct {
    // 优惠券实例ID/反核销回补ID
    VoucherId   string `json:"voucher_id,omitempty" xml:"voucher_id,omitempty"`
    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

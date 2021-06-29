package alsc

// PropertyRefundOpenInfo 
type PropertyRefundOpenInfo struct {
    // 回退积分是否成功
    PointSuccess   bool `json:"point_success,omitempty" xml:"point_success,omitempty"`
    // 回退储值是否成功
    ValueSuccess   bool `json:"value_success,omitempty" xml:"value_success,omitempty"`
    // 券实例状态
    VoucherStatusList   []VoucherStatusInfo `json:"voucher_status_list,omitempty" xml:"voucher_status_list>voucher_status_info,omitempty"`
}

package alsc

// PropertyVerifyOpenInfo 结构体
type PropertyVerifyOpenInfo struct {
	// 券码核销信息
	VoucherStatusList []VoucherStatusInfo `json:"voucher_status_list,omitempty" xml:"voucher_status_list>voucher_status_info,omitempty"`
	// 积分核销是否成功
	PointSuccess bool `json:"point_success,omitempty" xml:"point_success,omitempty"`
	// 储值核销是否成功
	ValueSuccess bool `json:"value_success,omitempty" xml:"value_success,omitempty"`
}

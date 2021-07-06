package alsc

// CustomerAggOpenInfo 结构体
type CustomerAggOpenInfo struct {
	// 卡模版列表
	CardOpenInfoList []CardOpenInfo `json:"card_open_info_list,omitempty" xml:"card_open_info_list>card_open_info,omitempty"`
	// 券列表
	CustomerVoucherFullOpenInfoList []CustomerVoucherFullOpenInfo `json:"customer_voucher_full_open_info_list,omitempty" xml:"customer_voucher_full_open_info_list>customer_voucher_full_open_info,omitempty"`
	// 顾客基础信息
	CustomerOpenInfo *CustomerOpenInfo `json:"customer_open_info,omitempty" xml:"customer_open_info,omitempty"`
	// 积分账户
	PointAccount *PointAccountOpenInfo `json:"point_account,omitempty" xml:"point_account,omitempty"`
}

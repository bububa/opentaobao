package alsc

// CustomerAggOpenInfo 
type CustomerAggOpenInfo struct {

    // 顾客基础信息
    CustomerOpenInfo   *CustomerOpenInfo `json:"customer_open_info,omitempty"`

    // 积分账户
    PointAccount   *PointAccountOpenInfo `json:"point_account,omitempty"`

    // 卡模版列表
    CardOpenInfoList   []CardOpenInfo `json:"card_open_info_list,omitempty"`

    // 券列表
    CustomerVoucherFullOpenInfoList   []CustomerVoucherFullOpenInfo `json:"customer_voucher_full_open_info_list,omitempty"`

}

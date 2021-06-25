package alsc

// CustomerGetOpenInfo 
type CustomerGetOpenInfo struct {

    // 卡实例信息
    CardOpenInfoList   []CardOpenInfo `json:"card_open_info_list,omitempty"`

    // 会员信息
    CustomerOpenInfo   *CustomerOpenInfo `json:"customer_open_info,omitempty"`

    // 积分信息
    PointAccountOpenInfo   *PointAccountOpenInfo `json:"point_account_open_info,omitempty"`

    // 券信息
    VoucherOpenInfoList   []VoucherOpenInfo `json:"voucher_open_info_list,omitempty"`

}

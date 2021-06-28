package alsc

// CustomerGetOpenInfo 
/* model for simplify = false
type CustomerGetOpenInfo struct {

    // 卡实例信息
    
    CardOpenInfoList  struct {
        CardOpenInfo  []CardOpenInfo `json:"card_open_info,omitempty"`
    } `json:"card_open_info_list,omitempty"`
    

    // 会员信息
    
    CustomerOpenInfo  *struct {
        CustomerOpenInfo  *CustomerOpenInfo `json:"customer_open_info,omitempty"`
    } `json:"customer_open_info,omitempty"`
    

    // 积分信息
    
    PointAccountOpenInfo  *struct {
        PointAccountOpenInfo  *PointAccountOpenInfo `json:"point_account_open_info,omitempty"`
    } `json:"point_account_open_info,omitempty"`
    

    // 券信息
    
    VoucherOpenInfoList  struct {
        VoucherOpenInfo  []VoucherOpenInfo `json:"voucher_open_info,omitempty"`
    } `json:"voucher_open_info_list,omitempty"`
    

}
*/

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

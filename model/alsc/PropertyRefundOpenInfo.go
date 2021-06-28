package alsc

// PropertyRefundOpenInfo 
/* model for simplify = false
type PropertyRefundOpenInfo struct {

    // 回退积分是否成功
    
    PointSuccess   bool `json:"point_success,omitempty"`
    

    // 回退储值是否成功
    
    ValueSuccess   bool `json:"value_success,omitempty"`
    

    // 券实例状态
    
    VoucherStatusList  struct {
        VoucherStatusInfo  []VoucherStatusInfo `json:"voucher_status_info,omitempty"`
    } `json:"voucher_status_list,omitempty"`
    

}
*/

// PropertyRefundOpenInfo 
type PropertyRefundOpenInfo struct {

    // 回退积分是否成功
    PointSuccess   bool `json:"point_success,omitempty"`

    // 回退储值是否成功
    ValueSuccess   bool `json:"value_success,omitempty"`

    // 券实例状态
    VoucherStatusList   []VoucherStatusInfo `json:"voucher_status_list,omitempty"`

}

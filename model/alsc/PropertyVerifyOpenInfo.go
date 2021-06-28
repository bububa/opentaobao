package alsc

// PropertyVerifyOpenInfo 
/* model for simplify = false
type PropertyVerifyOpenInfo struct {

    // 积分核销是否成功
    
    PointSuccess   bool `json:"point_success,omitempty"`
    

    // 储值核销是否成功
    
    ValueSuccess   bool `json:"value_success,omitempty"`
    

    // 券码核销信息
    
    VoucherStatusList  struct {
        VoucherStatusInfo  []VoucherStatusInfo `json:"voucher_status_info,omitempty"`
    } `json:"voucher_status_list,omitempty"`
    

}
*/

// PropertyVerifyOpenInfo 
type PropertyVerifyOpenInfo struct {

    // 积分核销是否成功
    PointSuccess   bool `json:"point_success,omitempty"`

    // 储值核销是否成功
    ValueSuccess   bool `json:"value_success,omitempty"`

    // 券码核销信息
    VoucherStatusList   []VoucherStatusInfo `json:"voucher_status_list,omitempty"`

}

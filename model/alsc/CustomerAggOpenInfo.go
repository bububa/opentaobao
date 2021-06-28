package alsc

// CustomerAggOpenInfo 
/* model for simplify = false
type CustomerAggOpenInfo struct {

    // 顾客基础信息
    
    CustomerOpenInfo  *struct {
        CustomerOpenInfo  *CustomerOpenInfo `json:"customer_open_info,omitempty"`
    } `json:"customer_open_info,omitempty"`
    

    // 积分账户
    
    PointAccount  *struct {
        PointAccountOpenInfo  *PointAccountOpenInfo `json:"point_account_open_info,omitempty"`
    } `json:"point_account,omitempty"`
    

    // 卡模版列表
    
    CardOpenInfoList  struct {
        CardOpenInfo  []CardOpenInfo `json:"card_open_info,omitempty"`
    } `json:"card_open_info_list,omitempty"`
    

    // 券列表
    
    CustomerVoucherFullOpenInfoList  struct {
        CustomerVoucherFullOpenInfo  []CustomerVoucherFullOpenInfo `json:"customer_voucher_full_open_info,omitempty"`
    } `json:"customer_voucher_full_open_info_list,omitempty"`
    

}
*/

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

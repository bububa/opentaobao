package alsc

// VoucherStatusInfo 
/* model for simplify = false
type VoucherStatusInfo struct {

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 券码
    
    VoucherId   string `json:"voucher_id,omitempty"`
    

}
*/

// VoucherStatusInfo 
type VoucherStatusInfo struct {

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 券码
    VoucherId   string `json:"voucher_id,omitempty"`

}

package alsc

// VoucherUsingOpenInfo 
/* model for simplify = false
type VoucherUsingOpenInfo struct {

    // 1
    
    VoucherStatusList  struct {
        VoucherStatus  []VoucherStatus `json:"voucher_status,omitempty"`
    } `json:"voucher_status_list,omitempty"`
    

}
*/

// VoucherUsingOpenInfo 
type VoucherUsingOpenInfo struct {

    // 1
    VoucherStatusList   []VoucherStatus `json:"voucher_status_list,omitempty"`

}

package mos

// CancelBillDto 
/* model for simplify = false
type CancelBillDto struct {

    // 取消备注
    
    CancelComments   string `json:"cancel_comments,omitempty"`
    

    // 结算行集合
    
    SettleLineNos  struct {
        String  []string `json:"string,omitempty"`
    } `json:"settle_line_nos,omitempty"`
    

    // 取消类型
    
    CancelType   string `json:"cancel_type,omitempty"`
    

    // 付款单号
    
    BillNo   string `json:"bill_no,omitempty"`
    

    // 备注
    
    ExtendParams   string `json:"extend_params,omitempty"`
    

}
*/

// CancelBillDto 
type CancelBillDto struct {

    // 取消备注
    CancelComments   string `json:"cancel_comments,omitempty"`

    // 结算行集合
    SettleLineNos   []string `json:"settle_line_nos,omitempty"`

    // 取消类型
    CancelType   string `json:"cancel_type,omitempty"`

    // 付款单号
    BillNo   string `json:"bill_no,omitempty"`

    // 备注
    ExtendParams   string `json:"extend_params,omitempty"`

}

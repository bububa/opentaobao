package mos

// CancelBillDto 
type CancelBillDto struct {

    // 取消备注
    CancelComments   string `json:"cancel_comments,omitempty"`

    // 结算行集合
    SettleLineNos   []String `json:"settle_line_nos,omitempty"`

    // 取消类型
    CancelType   string `json:"cancel_type,omitempty"`

    // 付款单号
    BillNo   string `json:"bill_no,omitempty"`

    // 备注
    ExtendParams   string `json:"extend_params,omitempty"`

}

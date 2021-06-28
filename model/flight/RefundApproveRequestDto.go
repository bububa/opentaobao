package flight

// RefundApproveRequestDto 
/* model for simplify = false
type RefundApproveRequestDto struct {

    // 申请单号
    
    ApplyId   string `json:"apply_id,omitempty"`
    

    // 国内国际标识
    
    DomesticIntl   int64 `json:"domestic_intl,omitempty"`
    

    // 退票数据
    
    RefundList  struct {
        RefundList  []RefundList `json:"refund_list,omitempty"`
    } `json:"refund_list,omitempty"`
    

    // 交易币种
    
    Currency   string `json:"currency,omitempty"`
    

}
*/

// RefundApproveRequestDto 
type RefundApproveRequestDto struct {

    // 申请单号
    ApplyId   string `json:"apply_id,omitempty"`

    // 国内国际标识
    DomesticIntl   int64 `json:"domestic_intl,omitempty"`

    // 退票数据
    RefundList   []RefundList `json:"refund_list,omitempty"`

    // 交易币种
    Currency   string `json:"currency,omitempty"`

}

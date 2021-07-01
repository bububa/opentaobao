package flight

// RefundApproveRequestDto 
type RefundApproveRequestDto struct {
    // 申请单号
    ApplyId   string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
    // 国内国际标识
    DomesticIntl   int64 `json:"domestic_intl,omitempty" xml:"domestic_intl,omitempty"`
    // 退票数据
    RefundList   []RefundList `json:"refund_list,omitempty" xml:"refund_list>refund_list,omitempty"`
    // 交易币种
    Currency   string `json:"currency,omitempty" xml:"currency,omitempty"`
}

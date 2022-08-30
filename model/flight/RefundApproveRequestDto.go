package flight

// RefundApproveRequestDto 结构体
type RefundApproveRequestDto struct {
	// 退票数据, 必填,
	RefundList []RefundList `json:"refund_list,omitempty" xml:"refund_list>refund_list,omitempty"`
	// 申请单号,必填
	ApplyId string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 交易币种: CNY:人民币, USD:美元, HKD:港元
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 国内国际标识:1:国内,2:国际
	DomesticIntl int64 `json:"domestic_intl,omitempty" xml:"domestic_intl,omitempty"`
}

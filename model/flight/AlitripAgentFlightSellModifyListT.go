package flight

// AlitripAgentFlightSellModifyListT 结构体
type AlitripAgentFlightSellModifyListT struct {
	// 申请单号
	ApplyId string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 飞猪订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 申请时间
	ApplyTime string `json:"apply_time,omitempty" xml:"apply_time,omitempty"`
	// 国内国际标识
	DomesticIntl int64 `json:"domestic_intl,omitempty" xml:"domestic_intl,omitempty"`
}

package flight

// AlitripAgentFlightSellTicketingListT 结构体
type AlitripAgentFlightSellTicketingListT struct {
	// 飞猪订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 国内国际标识
	DomesticIntl int64 `json:"domestic_intl,omitempty" xml:"domestic_intl,omitempty"`
}

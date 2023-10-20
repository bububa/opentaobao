package flight

import (
	"sync"
)

// AlitripAgentFlightSellTicketingListT 结构体
type AlitripAgentFlightSellTicketingListT struct {
	// 飞猪订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 服务时效
	ServeDeadline string `json:"serve_deadline,omitempty" xml:"serve_deadline,omitempty"`
	// 催出后服务时效
	UrgeServeDeadline string `json:"urge_serve_deadline,omitempty" xml:"urge_serve_deadline,omitempty"`
	// 2023-05-26 20:13:08后存在订单被取消的风险
	UrgeRiskDesc string `json:"urge_risk_desc,omitempty" xml:"urge_risk_desc,omitempty"`
	// 国内国际标识
	DomesticIntl int64 `json:"domestic_intl,omitempty" xml:"domestic_intl,omitempty"`
}

var poolAlitripAgentFlightSellTicketingListT = sync.Pool{
	New: func() any {
		return new(AlitripAgentFlightSellTicketingListT)
	},
}

// GetAlitripAgentFlightSellTicketingListT() 从对象池中获取AlitripAgentFlightSellTicketingListT
func GetAlitripAgentFlightSellTicketingListT() *AlitripAgentFlightSellTicketingListT {
	return poolAlitripAgentFlightSellTicketingListT.Get().(*AlitripAgentFlightSellTicketingListT)
}

// ReleaseAlitripAgentFlightSellTicketingListT 释放AlitripAgentFlightSellTicketingListT
func ReleaseAlitripAgentFlightSellTicketingListT(v *AlitripAgentFlightSellTicketingListT) {
	v.OrderId = ""
	v.PayTime = ""
	v.ServeDeadline = ""
	v.UrgeServeDeadline = ""
	v.UrgeRiskDesc = ""
	v.DomesticIntl = 0
	poolAlitripAgentFlightSellTicketingListT.Put(v)
}

package bus

import (
	"sync"
)

// AgentConfirmRefundRq 结构体
type AgentConfirmRefundRq struct {
	// 商家单号
	AgentOrderId string `json:"agent_order_id,omitempty" xml:"agent_order_id,omitempty"`
	// 按票退时商家退款票号
	AgentTicketId string `json:"agent_ticket_id,omitempty" xml:"agent_ticket_id,omitempty"`
	// 发车时间
	DepartDate string `json:"depart_date,omitempty" xml:"depart_date,omitempty"`
	// 退款乘客身份证号
	PassengerIdNum string `json:"passenger_id_num,omitempty" xml:"passenger_id_num,omitempty"`
	// 退款乘客手机号
	PassengerPhone string `json:"passenger_phone,omitempty" xml:"passenger_phone,omitempty"`
	// 退款金额，单位（元）
	RefundFee string `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 退款时间点
	RefundTime string `json:"refund_time,omitempty" xml:"refund_time,omitempty"`
	// 退款资金号唯一ID
	RefundTransId string `json:"refund_trans_id,omitempty" xml:"refund_trans_id,omitempty"`
	// 平台单号
	MainBizOrderId int64 `json:"main_biz_order_id,omitempty" xml:"main_biz_order_id,omitempty"`
	// 退款类型 0-退票
	RefundType int64 `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
}

var poolAgentConfirmRefundRq = sync.Pool{
	New: func() any {
		return new(AgentConfirmRefundRq)
	},
}

// GetAgentConfirmRefundRq() 从对象池中获取AgentConfirmRefundRq
func GetAgentConfirmRefundRq() *AgentConfirmRefundRq {
	return poolAgentConfirmRefundRq.Get().(*AgentConfirmRefundRq)
}

// ReleaseAgentConfirmRefundRq 释放AgentConfirmRefundRq
func ReleaseAgentConfirmRefundRq(v *AgentConfirmRefundRq) {
	v.AgentOrderId = ""
	v.AgentTicketId = ""
	v.DepartDate = ""
	v.PassengerIdNum = ""
	v.PassengerPhone = ""
	v.RefundFee = ""
	v.RefundTime = ""
	v.RefundTransId = ""
	v.MainBizOrderId = 0
	v.RefundType = 0
	poolAgentConfirmRefundRq.Put(v)
}

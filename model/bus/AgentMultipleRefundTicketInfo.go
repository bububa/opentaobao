package bus

import (
	"sync"
)

// AgentMultipleRefundTicketInfo 结构体
type AgentMultipleRefundTicketInfo struct {
	// 商家票号(唯一标识关联乘客)
	AgentTicketId string `json:"agent_ticket_id,omitempty" xml:"agent_ticket_id,omitempty"`
	// 退款人证件号(无票号时此参数必填)
	PassengerIdNum string `json:"passenger_id_num,omitempty" xml:"passenger_id_num,omitempty"`
	// 退票金额(分)
	RefundAmount int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
	// 退服务费金额(分)
	ServiceChargeRefundAmount int64 `json:"service_charge_refund_amount,omitempty" xml:"service_charge_refund_amount,omitempty"`
}

var poolAgentMultipleRefundTicketInfo = sync.Pool{
	New: func() any {
		return new(AgentMultipleRefundTicketInfo)
	},
}

// GetAgentMultipleRefundTicketInfo() 从对象池中获取AgentMultipleRefundTicketInfo
func GetAgentMultipleRefundTicketInfo() *AgentMultipleRefundTicketInfo {
	return poolAgentMultipleRefundTicketInfo.Get().(*AgentMultipleRefundTicketInfo)
}

// ReleaseAgentMultipleRefundTicketInfo 释放AgentMultipleRefundTicketInfo
func ReleaseAgentMultipleRefundTicketInfo(v *AgentMultipleRefundTicketInfo) {
	v.AgentTicketId = ""
	v.PassengerIdNum = ""
	v.RefundAmount = 0
	v.ServiceChargeRefundAmount = 0
	poolAgentMultipleRefundTicketInfo.Put(v)
}

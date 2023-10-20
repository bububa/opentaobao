package ship

import (
	"sync"
)

// ShipAgentConfirmRefundRq 结构体
type ShipAgentConfirmRefundRq struct {
	// 扩展属性
	AgentExtAttr string `json:"agent_ext_attr,omitempty" xml:"agent_ext_attr,omitempty"`
	// 商家订单id
	AgentOrderId string `json:"agent_order_id,omitempty" xml:"agent_order_id,omitempty"`
	// 退款资金号唯一ID
	AgentRefundTransId string `json:"agent_refund_trans_id,omitempty" xml:"agent_refund_trans_id,omitempty"`
	// offline:线下退票;online:线上退票
	AgentReturnMode string `json:"agent_return_mode,omitempty" xml:"agent_return_mode,omitempty"`
	// 退票时间
	AgentReturnTime string `json:"agent_return_time,omitempty" xml:"agent_return_time,omitempty"`
	// 按票的维度进行退票时，商家票号或者乘客Id必填一个,按单退填写订单号
	AgentTicketId string `json:"agent_ticket_id,omitempty" xml:"agent_ticket_id,omitempty"`
	// 乘客Id
	PassengerId string `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	// 退款金额(分)
	AgentRefundAmount int64 `json:"agent_refund_amount,omitempty" xml:"agent_refund_amount,omitempty"`
	// 退票结果编码,(退票失败时必填,参考标准错误码)
	AgentReturnTicketCode int64 `json:"agent_return_ticket_code,omitempty" xml:"agent_return_ticket_code,omitempty"`
	// 商家退票状态;1-成功,2-失败
	AgentReturnTicketStatus int64 `json:"agent_return_ticket_status,omitempty" xml:"agent_return_ticket_status,omitempty"`
	// 商家退票类型;0-按票退,1-按单退
	AgentReturnTicketType int64 `json:"agent_return_ticket_type,omitempty" xml:"agent_return_ticket_type,omitempty"`
	// 飞猪订单号
	AlitripOrderId int64 `json:"alitrip_order_id,omitempty" xml:"alitrip_order_id,omitempty"`
}

var poolShipAgentConfirmRefundRq = sync.Pool{
	New: func() any {
		return new(ShipAgentConfirmRefundRq)
	},
}

// GetShipAgentConfirmRefundRq() 从对象池中获取ShipAgentConfirmRefundRq
func GetShipAgentConfirmRefundRq() *ShipAgentConfirmRefundRq {
	return poolShipAgentConfirmRefundRq.Get().(*ShipAgentConfirmRefundRq)
}

// ReleaseShipAgentConfirmRefundRq 释放ShipAgentConfirmRefundRq
func ReleaseShipAgentConfirmRefundRq(v *ShipAgentConfirmRefundRq) {
	v.AgentExtAttr = ""
	v.AgentOrderId = ""
	v.AgentRefundTransId = ""
	v.AgentReturnMode = ""
	v.AgentReturnTime = ""
	v.AgentTicketId = ""
	v.PassengerId = ""
	v.AgentRefundAmount = 0
	v.AgentReturnTicketCode = 0
	v.AgentReturnTicketStatus = 0
	v.AgentReturnTicketType = 0
	v.AlitripOrderId = 0
	poolShipAgentConfirmRefundRq.Put(v)
}

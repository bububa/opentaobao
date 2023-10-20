package bus

import (
	"sync"
)

// SingleRefundInfo 结构体
type SingleRefundInfo struct {
	// 乘客身份证后四位
	CardNo string `json:"card_no,omitempty" xml:"card_no,omitempty"`
	// 退票价格
	RefundPrice string `json:"refund_price,omitempty" xml:"refund_price,omitempty"`
	// 乘客姓名
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 退票手续费
	RefundProcedurePrice string `json:"refund_procedure_price,omitempty" xml:"refund_procedure_price,omitempty"`
	// 退款详情
	RefundDetail string `json:"refund_detail,omitempty" xml:"refund_detail,omitempty"`
	// 代理商订单号
	AgentOrderId string `json:"agent_order_id,omitempty" xml:"agent_order_id,omitempty"`
	// 退票时间
	RefundTicketTime string `json:"refund_ticket_time,omitempty" xml:"refund_ticket_time,omitempty"`
	// 代理商票ID
	AgentTicketId string `json:"agent_ticket_id,omitempty" xml:"agent_ticket_id,omitempty"`
}

var poolSingleRefundInfo = sync.Pool{
	New: func() any {
		return new(SingleRefundInfo)
	},
}

// GetSingleRefundInfo() 从对象池中获取SingleRefundInfo
func GetSingleRefundInfo() *SingleRefundInfo {
	return poolSingleRefundInfo.Get().(*SingleRefundInfo)
}

// ReleaseSingleRefundInfo 释放SingleRefundInfo
func ReleaseSingleRefundInfo(v *SingleRefundInfo) {
	v.CardNo = ""
	v.RefundPrice = ""
	v.PassengerName = ""
	v.RefundProcedurePrice = ""
	v.RefundDetail = ""
	v.AgentOrderId = ""
	v.RefundTicketTime = ""
	v.AgentTicketId = ""
	poolSingleRefundInfo.Put(v)
}

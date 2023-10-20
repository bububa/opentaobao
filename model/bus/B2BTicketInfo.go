package bus

import (
	"sync"
)

// B2BTicketInfo 结构体
type B2BTicketInfo struct {
	// 乘客类型
	RiderCertType string `json:"rider_cert_type,omitempty" xml:"rider_cert_type,omitempty"`
	// 乘客姓名
	RiderName string `json:"rider_name,omitempty" xml:"rider_name,omitempty"`
	// 座位号
	RiderSeatNumber string `json:"rider_seat_number,omitempty" xml:"rider_seat_number,omitempty"`
	// 票ID
	TicketId string `json:"ticket_id,omitempty" xml:"ticket_id,omitempty"`
	// 退票手续费
	CommissionFee int64 `json:"commission_fee,omitempty" xml:"commission_fee,omitempty"`
	// 可退款金额
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 退票状态 1(新建立) 2（处理中）3（同意）4（拒绝）5（已退款）
	RefundStatus int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 服务费
	ServiceCharge int64 `json:"service_charge,omitempty" xml:"service_charge,omitempty"`
	// 子订单id
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 票价
	TicketPrice int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
}

var poolB2BTicketInfo = sync.Pool{
	New: func() any {
		return new(B2BTicketInfo)
	},
}

// GetB2BTicketInfo() 从对象池中获取B2BTicketInfo
func GetB2BTicketInfo() *B2BTicketInfo {
	return poolB2BTicketInfo.Get().(*B2BTicketInfo)
}

// ReleaseB2BTicketInfo 释放B2BTicketInfo
func ReleaseB2BTicketInfo(v *B2BTicketInfo) {
	v.RiderCertType = ""
	v.RiderName = ""
	v.RiderSeatNumber = ""
	v.TicketId = ""
	v.CommissionFee = 0
	v.RefundFee = 0
	v.RefundStatus = 0
	v.ServiceCharge = 0
	v.SubOrderId = 0
	v.TicketPrice = 0
	poolB2BTicketInfo.Put(v)
}

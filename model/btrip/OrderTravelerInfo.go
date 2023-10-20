package btrip

import (
	"sync"
)

// OrderTravelerInfo 结构体
type OrderTravelerInfo struct {
	// 乘客名称
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 乘客类型
	PassengerType string `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
	// 票号
	TicketNo string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// 用户编号
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// openTicket状态
	OpenTicketStatus int64 `json:"open_ticket_status,omitempty" xml:"open_ticket_status,omitempty"`
}

var poolOrderTravelerInfo = sync.Pool{
	New: func() any {
		return new(OrderTravelerInfo)
	},
}

// GetOrderTravelerInfo() 从对象池中获取OrderTravelerInfo
func GetOrderTravelerInfo() *OrderTravelerInfo {
	return poolOrderTravelerInfo.Get().(*OrderTravelerInfo)
}

// ReleaseOrderTravelerInfo 释放OrderTravelerInfo
func ReleaseOrderTravelerInfo(v *OrderTravelerInfo) {
	v.PassengerName = ""
	v.PassengerType = ""
	v.TicketNo = ""
	v.UserId = ""
	v.OpenTicketStatus = 0
	poolOrderTravelerInfo.Put(v)
}
